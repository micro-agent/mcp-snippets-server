package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/micro-agent/micro-agent-go/agent/helpers"
	"github.com/micro-agent/micro-agent-go/agent/mu"
	"github.com/micro-agent/micro-agent-go/agent/rag"
	"github.com/openai/openai-go/v2" // imported as openai
	"github.com/openai/openai-go/v2/option"
)

var store rag.MemoryVectorStore
var embeddingAgent mu.Agent
var embeddingsModel string

// determineDelimiter returns an appropriate delimiter based on min/max constraints
func determineDelimiter(minDelimiter, maxDelimiter string) string {
	// If min is longer than max, use max
	if len(minDelimiter) > len(maxDelimiter) {
		return maxDelimiter
	}
	
	// Check if max delimiter length meets minimum requirement
	if len(maxDelimiter) >= len(minDelimiter) {
		// Use a delimiter that's between min and max length
		// For simplicity, we'll use the min delimiter as it's guaranteed to be valid
		return minDelimiter
	}
	
	// Fallback to a default delimiter
	return "----------"
}

func main() {
	ctx := context.Background()

	// Create MCP server
	s := server.NewMCPServer(
		"mcp-snippets-server",
		"0.0.0",
	)

	baseURL := helpers.GetEnvOrDefault("MODEL_RUNNER_BASE_URL", "http://localhost:12434/engines/llama.cpp/v1/")
	embeddingsModel = helpers.GetEnvOrDefault("EMBEDDING_MODEL", "ai/mxbai-embed-large:latest")
	jsonStoreFilePath := helpers.GetEnvOrDefault("JSON_STORE_FILE_PATH", "rag-memory-store.json")
	minDelimiter := helpers.GetEnvOrDefault("MINIMUM_DELIMITER", "----------")
	maxDelimiter := helpers.GetEnvOrDefault("MAXIMUM_DELIMITER", "----------------------------------------")

	client := openai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey(""),
	)

	// EMBEDDING AGENT: Create an embedding agent to generate embeddings
	var err error
	embeddingAgent, err = mu.NewAgent(ctx, "vector-agent",
		mu.WithClient(client),
		mu.WithEmbeddingParams(
			openai.EmbeddingNewParams{
				Model: embeddingsModel,
			},
		),
	)
	if err != nil {
		fmt.Println("🔶 Error creating embedding agent", err)
		panic(err)
	}

	// -------------------------------------------------
	// Create a vector store
	// -------------------------------------------------
	store = rag.MemoryVectorStore{
		Records: make(map[string]rag.VectorRecord),
	}

	// Load the vector store from a file if it exists
	err = store.Load(jsonStoreFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("🚀 No existing vector store found, starting fresh.")

			// =================================================
			// CHUNKS:
			// =================================================
			contents, err := helpers.GetContentFiles(".", ".md")
			if err != nil {
				log.Fatalln("😡 Error getting content files:", err)
			}
			chunks := []string{}
			fmt.Println("💡 Found", len(contents), "content files to process.")
			//fmt.Println("📂 Processing content files...", contents)
			fmt.Println("📝 Processing(Chunking) content files...")

			for _, content := range contents {
				// Determine appropriate delimiter based on min/max constraints
				delimiter := determineDelimiter(minDelimiter, maxDelimiter)
				fmt.Println("📏 Using delimiter:", delimiter, "(length:", len(delimiter), ")")
				chunks = append(chunks, rag.SplitTextWithDelimiter(content, delimiter)...)
			}

			// -------------------------------------------------
			// Create and save the embeddings from the chunks
			// -------------------------------------------------
			fmt.Println("⏳ Creating the embeddings...")

			for idx, chunk := range chunks {

				fmt.Println("🔶 Chunk", idx, ":", chunk)
				embeddingVector, err := embeddingAgent.GenerateEmbeddingVector(chunk)

				if err != nil {
					fmt.Println(err)
					fmt.Println(chunk)
				} else {
					_, errSave := store.Save(rag.VectorRecord{
						Prompt:    chunk,
						Embedding: embeddingVector,
					})
					if errSave != nil {
						fmt.Println("😡:", errSave)
					}
					fmt.Println("✅ Chunk", idx, "saved with embedding:", len(embeddingVector))
				}
			}

			fmt.Println("✋", "Embeddings created, total of records", len(store.Records))
			err = store.Persist(jsonStoreFilePath)
			if err != nil {
				log.Fatalln("😡 Error saving vector store:", err)
			}
			fmt.Println("✅ Vector store saved to", jsonStoreFilePath)
			fmt.Println("💾 Vector store initialized with", len(store.Records), "records.")
			fmt.Println()

		} else {
			log.Fatalln("Error loading vector store:", err)
		}
	} else {
		log.Println("Vector store loaded successfully, total records:", len(store.Records))
	}

	// =================================================
	// TOOLS:
	// =================================================
	searchInDoc := mcp.NewTool("search_snippet",
		mcp.WithDescription(`Find one or more snippets related to the topic.`),
		mcp.WithString("topic",
			mcp.Required(),
			mcp.Description("Search topic or question to find relevant snippets."),
		),
	)
	s.AddTool(searchInDoc, searchInDocHandler)

	// Start the HTTP server
	httpPort := os.Getenv("MCP_HTTP_PORT")
	fmt.Println("🌍 MCP HTTP Port:", httpPort)
	if httpPort == "" {
		httpPort = "9090"
	}

	log.Println("MCP StreamableHTTP server is running on port", httpPort)

	// Create a custom mux to handle both MCP and health endpoints
	mux := http.NewServeMux()

	// Add healthcheck endpoint
	mux.HandleFunc("/health", healthCheckHandler)

	// Add MCP endpoint
	httpServer := server.NewStreamableHTTPServer(s,
		server.WithEndpointPath("/mcp"),
	)

	// Register MCP handler with the mux
	mux.Handle("/mcp", httpServer)

	// Start the HTTP server with custom mux
	log.Fatal(http.ListenAndServe(":"+httpPort, mux))
}

func searchInDocHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	args := request.GetArguments()
	topicArg, exists := args["topic"]
	if !exists || topicArg == nil {
		return nil, fmt.Errorf("missing required parameter 'topic'")
	}
	userQuestion, ok := topicArg.(string)
	if !ok {
		return nil, fmt.Errorf("parameter 'topic' must be a string")
	}

	fmt.Println("🔍 Searching for question:", userQuestion)

	// -------------------------------------------------
	// Search for similarities
	// -------------------------------------------------

	fmt.Println("⏳ Searching for similarities...")

	// -------------------------------------------------
	// Create embedding from the user question
	// -------------------------------------------------
	questionEmbeddingVector, err := embeddingAgent.GenerateEmbeddingVector(userQuestion)
	if err != nil {
		log.Fatal("😡:", err)
	}
	// -------------------------------------------------
	// Create a vector record from the user embedding
	// -------------------------------------------------
	questionRecord := rag.VectorRecord{Embedding: questionEmbeddingVector}



	threshold := helpers.StringToFloat(helpers.GetEnvOrDefault("LIMIT", "0.6"))
	topN := helpers.StringToInt(helpers.GetEnvOrDefault("MAX_RESULTS", "2"))

	similarities, err := store.SearchTopNSimilarities(questionRecord, threshold, topN)
	if err != nil {
		log.Fatal("😡:", err)
	}

	documentsContent := "Documents:\n"

	for _, similarity := range similarities {
		fmt.Println("✅ CosineSimilarity:", similarity.CosineSimilarity, "Chunk:", similarity.Prompt)
		documentsContent += similarity.Prompt
	}
	documentsContent += "\n"
	fmt.Println("✋", "Similarities found, total of records", len(similarities))
	fmt.Println()

	return mcp.NewToolResultText(documentsContent), nil
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if vector store is initialized and has records
	if len(store.Records) == 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
		response := map[string]interface{}{
			"status": "unhealthy",
			"reason": "vector store not initialized",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]any{
		"status":           "healthy",
		"records":          len(store.Records),
		"embeddings_model": embeddingsModel,
	}
	json.NewEncoder(w).Encode(response)
}
