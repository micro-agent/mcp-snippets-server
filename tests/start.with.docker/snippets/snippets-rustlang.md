## Hello World
Basic program structure and main function
```rust
fn main() {
    println!("Hello, World!");
}
```

----------

## Variable Declaration
Different ways to declare and initialize variables
```rust
fn main() {
    let name = "John";           // Immutable
    let mut age = 30;           // Mutable
    let city: &str = "New York"; // With type annotation
    
    age += 1; // Modify mutable variable
    
    println!("Name: {}, Age: {}, City: {}", name, age, city);
}
```

----------

## Constants and Static
Declaring constants and static variables
```rust
const PI: f64 = 3.14159;
const MAX_USERS: u32 = 100;
static GREETING: &str = "Hello";

fn main() {
    println!("Pi: {}, Max Users: {}, Greeting: {}", PI, MAX_USERS, GREETING);
}
```

----------

## Data Types
Working with basic data types
```rust
fn main() {
    let integer: i32 = 42;
    let float: f64 = 3.14;
    let boolean: bool = true;
    let character: char = '🦀';
    let string: String = String::from("Rust");
    let str_slice: &str = "slice";
    
    println!("int: {}, float: {}, bool: {}, char: {}", integer, float, boolean, character);
    println!("String: {}, &str: {}", string, str_slice);
}
```

----------

## Arrays and Vectors
Working with arrays and vectors
```rust
fn main() {
    let arr: [i32; 3] = [1, 2, 3];
    let mut vec = vec![1, 2, 3, 4];
    
    vec.push(5);
    vec.pop();
    
    println!("Array: {:?}", arr);
    println!("Vector: {:?}, len: {}", vec, vec.len());
    
    for (i, value) in vec.iter().enumerate() {
        println!("Index {}: {}", i, value);
    }
}
```

----------

## Hash Maps
Creating and manipulating hash maps
```rust
use std::collections::HashMap;

fn main() {
    let mut map = HashMap::new();
    map.insert("apple", 5);
    map.insert("banana", 3);
    
    // Check if key exists
    match map.get("apple") {
        Some(value) => println!("Apple: {}", value),
        None => println!("Apple not found"),
    }
    
    for (key, value) in &map {
        println!("{}: {}", key, value);
    }
}
```

----------

## Control Flow - If/Else
Conditional statements
```rust
fn main() {
    let age = 25;
    
    if age >= 18 {
        println!("Adult");
    } else {
        println!("Minor");
    }
    
    // If as expression
    let status = if age >= 18 { "adult" } else { "minor" };
    println!("Status: {}", status);
    
    let score = 85;
    let grade = if score >= 90 { "A" } else if score >= 80 { "B" } else { "C" };
    println!("Grade: {}", grade);
}
```

----------

## Loops
Different types of loops
```rust
fn main() {
    // Loop with break
    let mut counter = 0;
    loop {
        if counter >= 3 {
            break;
        }
        println!("Loop: {}", counter);
        counter += 1;
    }
    
    // While loop
    let mut n = 3;
    while n > 0 {
        println!("While: {}", n);
        n -= 1;
    }
    
    // For loop
    for i in 0..3 {
        println!("For: {}", i);
    }
}
```

----------

## Match Statements
Pattern matching with match
```rust
fn main() {
    let number = 3;
    
    match number {
        1 => println!("One"),
        2 | 3 => println!("Two or Three"),
        4..=10 => println!("Four to Ten"),
        _ => println!("Something else"),
    }
    
    let option_value = Some(5);
    match option_value {
        Some(x) => println!("Got value: {}", x),
        None => println!("No value"),
    }
}
```

----------

## Functions
Function declaration and parameters
```rust
fn add(a: i32, b: i32) -> i32 {
    a + b // No semicolon = return value
}

fn divide(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        Err("Division by zero".to_string())
    } else {
        Ok(a / b)
    }
}

fn main() {
    let sum = add(5, 3);
    println!("Sum: {}", sum);
    
    match divide(10.0, 2.0) {
        Ok(result) => println!("Division: {}", result),
        Err(e) => println!("Error: {}", e),
    }
}
```

----------

## Structs
Defining and using structs
```rust
struct Person {
    name: String,
    age: u32,
}

impl Person {
    fn new(name: String, age: u32) -> Person {
        Person { name, age }
    }
    
    fn greet(&self) -> String {
        format!("Hello, I'm {}", self.name)
    }
}

fn main() {
    let person = Person::new("Alice".to_string(), 30);
    println!("Person: {} ({})", person.name, person.age);
    println!("{}", person.greet());
}
```

----------

## Enums
Defining and using enums
```rust
enum Color {
    Red,
    Green,
    Blue,
    RGB(u8, u8, u8),
}

fn main() {
    let red = Color::Red;
    let custom = Color::RGB(255, 128, 0);
    
    match red {
        Color::Red => println!("It's red!"),
        Color::Green => println!("It's green!"),
        Color::Blue => println!("It's blue!"),
        Color::RGB(r, g, b) => println!("RGB: {}, {}, {}", r, g, b),
    }
}
```

----------

## Option and Result
Working with Option and Result types
```rust
fn find_user(id: u32) -> Option<String> {
    if id == 1 {
        Some("Alice".to_string())
    } else {
        None
    }
}

fn parse_number(s: &str) -> Result<i32, std::num::ParseIntError> {
    s.parse()
}

fn main() {
    match find_user(1) {
        Some(name) => println!("Found user: {}", name),
        None => println!("User not found"),
    }
    
    match parse_number("42") {
        Ok(n) => println!("Parsed: {}", n),
        Err(e) => println!("Parse error: {}", e),
    }
}
```

----------

## Error Handling
Proper error handling with ? operator
```rust
use std::fs;
use std::io;

fn read_file_content(path: &str) -> Result<String, io::Error> {
    let content = fs::read_to_string(path)?;
    Ok(content.trim().to_string())
}

fn main() {
    match read_file_content("test.txt") {
        Ok(content) => println!("File content: {}", content),
        Err(e) => println!("Error reading file: {}", e),
    }
}
```

----------

## Ownership and Borrowing
Understanding ownership rules
```rust
fn take_ownership(s: String) {
    println!("Took ownership of: {}", s);
}

fn borrow_string(s: &String) {
    println!("Borrowed: {}", s);
}

fn main() {
    let s1 = String::from("Hello");
    borrow_string(&s1); // Borrow
    println!("Still have: {}", s1);
    
    take_ownership(s1); // Move
    // println!("{}", s1); // Would cause compile error
}
```

----------

## Traits
Defining and implementing traits
```rust
trait Drawable {
    fn draw(&self);
}

struct Circle {
    radius: f64,
}

struct Rectangle {
    width: f64,
    height: f64,
}

impl Drawable for Circle {
    fn draw(&self) {
        println!("Drawing circle with radius {}", self.radius);
    }
}

impl Drawable for Rectangle {
    fn draw(&self) {
        println!("Drawing rectangle {}x{}", self.width, self.height);
    }
}

fn main() {
    let circle = Circle { radius: 5.0 };
    let rect = Rectangle { width: 3.0, height: 4.0 };
    
    circle.draw();
    rect.draw();
}
```

----------

## Generics
Using generic types and functions
```rust
fn find_largest<T: PartialOrd + Copy>(list: &[T]) -> T {
    let mut largest = list[0];
    for &item in list {
        if item > largest {
            largest = item;
        }
    }
    largest
}

struct Point<T> {
    x: T,
    y: T,
}

fn main() {
    let numbers = vec![34, 50, 25, 100, 65];
    let largest = find_largest(&numbers);
    println!("Largest number: {}", largest);
    
    let point = Point { x: 5, y: 10 };
    println!("Point: ({}, {})", point.x, point.y);
}
```

----------

## Lifetimes
Working with lifetime annotations
```rust
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

struct ImportantExcerpt<'a> {
    part: &'a str,
}

fn main() {
    let string1 = String::from("long string");
    let string2 = "short";
    
    let result = longest(&string1, string2);
    println!("Longest: {}", result);
    
    let excerpt = ImportantExcerpt { part: &string1 };
    println!("Excerpt: {}", excerpt.part);
}
```

----------

## Closures
Using closures and functional programming
```rust
fn main() {
    let numbers = vec![1, 2, 3, 4, 5];
    
    // Closure that captures environment
    let multiplier = 2;
    let doubled: Vec<i32> = numbers.iter().map(|x| x * multiplier).collect();
    println!("Doubled: {:?}", doubled);
    
    // Filter and collect
    let evens: Vec<i32> = numbers.into_iter().filter(|&x| x % 2 == 0).collect();
    println!("Evens: {:?}", evens);
    
    // Using closure as parameter
    let add_one = |x| x + 1;
    println!("5 + 1 = {}", add_one(5));
}
```

----------

## Iterators
Working with iterators
```rust
fn main() {
    let vec = vec![1, 2, 3, 4, 5];
    
    // Iterator methods
    let sum: i32 = vec.iter().sum();
    let max = vec.iter().max();
    println!("Sum: {}, Max: {:?}", sum, max);
    
    // Chain operations
    let result: Vec<i32> = vec
        .iter()
        .filter(|&&x| x > 2)
        .map(|x| x * 2)
        .collect();
    println!("Filtered and doubled: {:?}", result);
    
    // Custom iterator
    for (i, value) in vec.iter().enumerate() {
        println!("Index {}: {}", i, value);
    }
}
```

----------

## File I/O
Reading and writing files
```rust
use std::fs;
use std::io::Write;

fn main() -> std::io::Result<()> {
    // Write to file
    let content = "Hello, Rust!";
    fs::write("test.txt", content)?;
    
    // Read from file
    let read_content = fs::read_to_string("test.txt")?;
    println!("File content: {}", read_content);
    
    // Append to file
    let mut file = std::fs::OpenOptions::new()
        .append(true)
        .open("test.txt")?;
    writeln!(file, "Appended line")?;
    
    // Clean up
    fs::remove_file("test.txt")?;
    Ok(())
}
```

----------

## JSON Handling
Serialize and deserialize JSON with serde
```rust
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
struct Person {
    name: String,
    age: u32,
    email: Option<String>,
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let person = Person {
        name: "Alice".to_string(),
        age: 30,
        email: Some("alice@example.com".to_string()),
    };
    
    // Serialize to JSON
    let json = serde_json::to_string(&person)?;
    println!("JSON: {}", json);
    
    // Deserialize from JSON
    let parsed: Person = serde_json::from_str(&json)?;
    println!("Parsed: {:?}", parsed);
    
    Ok(())
}
```

----------

## HTTP Client
Making HTTP requests with reqwest
```rust
use reqwest;
use tokio;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let response = reqwest::get("https://httpbin.org/json").await?;
    let status = response.status();
    let body = response.text().await?;
    
    println!("Status: {}", status);
    println!("Body: {}", body);
    
    // POST request with JSON
    let client = reqwest::Client::new();
    let json_body = serde_json::json!({
        "name": "John",
        "age": 30
    });
    
    let post_response = client
        .post("https://httpbin.org/post")
        .json(&json_body)
        .send()
        .await?;
    
    println!("POST Status: {}", post_response.status());
    Ok(())
}
```

----------

## HTTP Server
Creating a simple HTTP server with warp
```rust
use warp::Filter;

#[tokio::main]
async fn main() {
    let hello = warp::path!("hello" / String)
        .map(|name| format!("Hello, {}!", name));
    
    let json_route = warp::path("json")
        .map(|| warp::reply::json(&serde_json::json!({
            "message": "Hello, JSON!"
        })));
    
    let routes = hello.or(json_route);
    
    println!("Server starting on localhost:3030");
    warp::serve(routes)
        .run(([127, 0, 0, 1], 3030))
        .await;
}
```

----------

## String Manipulation
Common string operations
```rust
fn main() {
    let text = "  Hello, World!  ";
    
    println!("Original: '{}'", text);
    println!("Trimmed: '{}'", text.trim());
    println!("Uppercase: {}", text.to_uppercase());
    println!("Contains 'World': {}", text.contains("World"));
    
    let words: Vec<&str> = text.trim().split_whitespace().collect();
    println!("Words: {:?}", words);
    
    // String manipulation
    let mut s = String::from("Hello");
    s.push_str(", World!");
    s.push('!');
    println!("Built string: {}", s);
    
    // Parse string to number
    let num: i32 = "42".parse().unwrap_or(0);
    println!("Parsed number: {}", num);
}
```

----------

## Date and Time
Working with chrono for date/time operations
```rust
use chrono::{DateTime, Utc, Local, Duration};

fn main() {
    let now: DateTime<Utc> = Utc::now();
    let local_now: DateTime<Local> = Local::now();
    
    println!("UTC now: {}", now);
    println!("Local now: {}", local_now);
    println!("Formatted: {}", now.format("%Y-%m-%d %H:%M:%S"));
    
    // Date arithmetic
    let tomorrow = now + Duration::days(1);
    let one_hour_ago = now - Duration::hours(1);
    
    println!("Tomorrow: {}", tomorrow);
    println!("One hour ago: {}", one_hour_ago);
    
    // Parse date string
    let date_str = "2023-12-25T10:30:00Z";
    match DateTime::parse_from_rfc3339(date_str) {
        Ok(parsed) => println!("Parsed: {}", parsed),
        Err(e) => println!("Parse error: {}", e),
    }
}
```

----------

## Testing
Unit tests and integration tests
```rust
fn add(a: i32, b: i32) -> i32 {
    a + b
}

fn divide(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        Err("Division by zero".to_string())
    } else {
        Ok(a / b)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        assert_eq!(add(2, 3), 5);
        assert_eq!(add(-1, 1), 0);
    }

    #[test]
    fn test_divide() {
        assert_eq!(divide(10.0, 2.0), Ok(5.0));
        assert!(divide(10.0, 0.0).is_err());
    }

    #[test]
    #[should_panic]
    fn test_panic() {
        panic!("This test should panic");
    }
}

fn main() {
    println!("Run with: cargo test");
}
```

----------

## Async Programming
Asynchronous programming with async/await
```rust
use tokio::time::{sleep, Duration};

async fn fetch_data(id: u32) -> String {
    sleep(Duration::from_millis(100)).await;
    format!("Data for ID: {}", id)
}

async fn process_multiple() {
    let mut handles = vec![];
    
    for i in 1..=3 {
        let handle = tokio::spawn(async move {
            fetch_data(i).await
        });
        handles.push(handle);
    }
    
    for handle in handles {
        match handle.await {
            Ok(result) => println!("{}", result),
            Err(e) => println!("Error: {}", e),
        }
    }
}

#[tokio::main]
async fn main() {
    println!("Starting async operations...");
    process_multiple().await;
    println!("All operations completed");
}
```

----------

## Threads
Multi-threading and synchronization
```rust
use std::thread;
use std::sync::{Arc, Mutex};
use std::time::Duration;

fn main() {
    // Simple thread
    let handle = thread::spawn(|| {
        for i in 1..=5 {
            println!("Thread: {}", i);
            thread::sleep(Duration::from_millis(100));
        }
    });
    
    // Main thread work
    for i in 1..=3 {
        println!("Main: {}", i);
        thread::sleep(Duration::from_millis(150));
    }
    
    handle.join().unwrap();
    
    // Shared state with Arc<Mutex<T>>
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];
    
    for _ in 0..3 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();
            *num += 1;
        });
        handles.push(handle);
    }
    
    for handle in handles {
        handle.join().unwrap();
    }
    
    println!("Final counter: {}", *counter.lock().unwrap());
}
```

----------

## Command Line Arguments
Parsing command line arguments with clap
```rust
use clap::{Arg, App};

fn main() {
    let matches = App::new("My App")
        .version("1.0")
        .author("Your Name")
        .about("Does awesome things")
        .arg(Arg::with_name("name")
            .short("n")
            .long("name")
            .value_name("NAME")
            .help("Sets a custom name")
            .takes_value(true)
            .default_value("World"))
        .arg(Arg::with_name("verbose")
            .short("v")
            .long("verbose")
            .help("Enable verbose output"))
        .get_matches();
    
    let name = matches.value_of("name").unwrap();
    let verbose = matches.is_present("verbose");
    
    if verbose {
        println!("Verbose mode enabled");
        println!("Name parameter: {}", name);
    }
    
    println!("Hello, {}!", name);
}
```

----------

## Regular Expressions
Pattern matching with regex
```rust
use regex::Regex;

fn main() {
    let email_regex = Regex::new(r"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$").unwrap();
    
    let emails = vec!["user@example.com", "invalid.email", "test@domain.co.uk"];
    
    for email in emails {
        if email_regex.is_match(email) {
            println!("{}: Valid", email);
        } else {
            println!("{}: Invalid", email);
        }
    }
    
    // Find all matches
    let text = "Contact us at info@company.com or support@company.com";
    let matches: Vec<&str> = email_regex.find_iter(text).map(|m| m.as_str()).collect();
    println!("Found emails: {:?}", matches);
    
    // Replace matches
    let replaced = email_regex.replace_all(text, "[EMAIL]");
    println!("Replaced: {}", replaced);
}
```

----------

## Logging
Structured logging with log and env_logger
```rust
use log::{info, warn, error, debug};

fn main() {
    env_logger::init();
    
    debug!("This is a debug message");
    info!("Application started");
    warn!("This is a warning");
    error!("This is an error message");
    
    // Structured logging
    info!("User logged in"; "user_id" => 123, "ip" => "192.168.1.1");
    
    // Log with formatting
    let user_name = "Alice";
    let login_count = 5;
    info!("User {} has logged in {} times", user_name, login_count);
}
```

----------

## Configuration
Managing configuration with config crate
```rust
use serde::Deserialize;

#[derive(Debug, Deserialize)]
struct ServerConfig {
    host: String,
    port: u16,
}

#[derive(Debug, Deserialize)]
struct DatabaseConfig {
    url: String,
    max_connections: u32,
}

#[derive(Debug, Deserialize)]
struct AppConfig {
    server: ServerConfig,
    database: DatabaseConfig,
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut settings = config::Config::default();
    
    // Add configuration sources
    settings
        .merge(config::File::with_name("config"))?
        .merge(config::Environment::with_prefix("APP"))?;
    
    // Deserialize configuration
    let app_config: AppConfig = settings.try_into()?;
    
    println!("Config: {:?}", app_config);
    println!("Server will run on {}:{}", app_config.server.host, app_config.server.port);
    
    Ok(())
}
```