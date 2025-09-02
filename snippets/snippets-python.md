## Hello World
Basic program structure and main function
```python
def main():
    print("Hello, World!")

if __name__ == "__main__":
    main()
```

----------

## Variable Declaration
Different ways to declare and initialize variables
```python
name = "John"           # String
age = 30               # Integer
height = 5.9           # Float
is_active = True       # Boolean
fruits = ["apple", "banana"]  # List
x, y = 10, 20          # Multiple assignment

print(f"Name: {name}, Age: {age}, Height: {height}")
```

----------

## Data Types
Working with basic data types
```python
# Numbers
integer = 42
float_num = 3.14
complex_num = 1 + 2j

# Strings
string = "Python"
multiline = """This is a
multiline string"""

# Collections
list_example = [1, 2, 3, 4]
tuple_example = (1, 2, 3)
set_example = {1, 2, 3}
dict_example = {"name": "Alice", "age": 30}

print(f"Types: {type(integer)}, {type(string)}, {type(list_example)}")
```

----------

## Lists and List Comprehensions
Working with lists and comprehensions
```python
numbers = [1, 2, 3, 4, 5]
names = ["Alice", "Bob", "Charlie"]

# List operations
numbers.append(6)
numbers.insert(0, 0)
first = numbers.pop(0)

# List comprehensions
squares = [x**2 for x in numbers]
evens = [x for x in numbers if x % 2 == 0]
pairs = [(x, y) for x in range(3) for y in range(3)]

print(f"Squares: {squares}")
print(f"Evens: {evens}")
```

----------

## Dictionaries
Creating and manipulating dictionaries
```python
person = {"name": "Alice", "age": 30, "city": "New York"}

# Dictionary operations
person["email"] = "alice@example.com"
age = person.get("age", 0)
keys = list(person.keys())
values = list(person.values())

# Dictionary comprehension
squares_dict = {x: x**2 for x in range(5)}

# Iterating
for key, value in person.items():
    print(f"{key}: {value}")
```

----------

## Control Flow - If/Else
Conditional statements
```python
age = 25
score = 85

if age >= 18:
    print("Adult")
elif age >= 13:
    print("Teenager")
else:
    print("Child")

# Ternary operator
status = "adult" if age >= 18 else "minor"

# Multiple conditions
if 80 <= score <= 100:
    grade = "A"
elif 70 <= score < 80:
    grade = "B"
else:
    grade = "C"

print(f"Status: {status}, Grade: {grade}")
```

----------

## Loops
Different types of loops
```python
# For loop with range
for i in range(5):
    print(f"Number: {i}")

# For loop with list
fruits = ["apple", "banana", "cherry"]
for fruit in fruits:
    print(f"Fruit: {fruit}")

# For loop with enumerate
for index, fruit in enumerate(fruits):
    print(f"{index}: {fruit}")

# While loop
count = 0
while count < 3:
    print(f"Count: {count}")
    count += 1

# Loop with else
for i in range(3):
    print(i)
else:
    print("Loop completed")
```

----------

## Functions
Function definition and parameters
```python
def greet(name, greeting="Hello"):
    return f"{greeting}, {name}!"

def add_numbers(*args):
    return sum(args)

def create_profile(**kwargs):
    return {k: v for k, v in kwargs.items()}

# Lambda functions
square = lambda x: x**2
add = lambda x, y: x + y

# Function calls
message = greet("Alice")
total = add_numbers(1, 2, 3, 4, 5)
profile = create_profile(name="Bob", age=25, city="Boston")

print(f"Message: {message}")
print(f"Total: {total}")
print(f"Square of 5: {square(5)}")
```

----------

## Classes and Objects
Object-oriented programming
```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
        self._id = id(self)  # Private-ish attribute

    def greet(self):
        return f"Hello, I'm {self.name}"

    def __str__(self):
        return f"Person(name={self.name}, age={self.age})"

    @classmethod
    def from_string(cls, person_str):
        name, age = person_str.split('-')
        return cls(name, int(age))

    @staticmethod
    def is_adult(age):
        return age >= 18

# Usage
person = Person("Alice", 30)
print(person.greet())
print(f"Is adult: {Person.is_adult(person.age)}")

person2 = Person.from_string("Bob-25")
print(person2)
```

----------

## Inheritance
Class inheritance and method overriding
```python
class Animal:
    def __init__(self, name):
        self.name = name

    def speak(self):
        pass

    def info(self):
        return f"This is {self.name}"

class Dog(Animal):
    def __init__(self, name, breed):
        super().__init__(name)
        self.breed = breed

    def speak(self):
        return "Woof!"

    def info(self):
        return f"{super().info()}, a {self.breed} dog"

class Cat(Animal):
    def speak(self):
        return "Meow!"

# Usage
dog = Dog("Buddy", "Golden Retriever")
cat = Cat("Whiskers")

print(f"{dog.name} says: {dog.speak()}")
print(dog.info())
print(f"{cat.name} says: {cat.speak()}")
```

----------

## Exception Handling
Try-catch blocks and custom exceptions
```python
class CustomError(Exception):
    def __init__(self, message):
        self.message = message
        super().__init__(self.message)

def divide_numbers(a, b):
    if b == 0:
        raise CustomError("Cannot divide by zero")
    return a / b

def safe_conversion(value):
    try:
        result = int(value)
        return result
    except ValueError:
        print(f"Cannot convert '{value}' to integer")
        return None
    except Exception as e:
        print(f"Unexpected error: {e}")
        return None
    finally:
        print("Conversion attempt completed")

# Usage
try:
    result = divide_numbers(10, 2)
    print(f"Result: {result}")
    
    bad_result = divide_numbers(10, 0)
except CustomError as e:
    print(f"Custom error: {e.message}")

number = safe_conversion("123")
invalid = safe_conversion("abc")
```

----------

## File I/O
Reading and writing files
```python
import os

# Writing to file
content = "Hello, Python!\nSecond line"
with open("test.txt", "w") as file:
    file.write(content)

# Reading from file
with open("test.txt", "r") as file:
    file_content = file.read()
    print(f"File content:\n{file_content}")

# Reading lines
with open("test.txt", "r") as file:
    lines = file.readlines()
    for i, line in enumerate(lines, 1):
        print(f"Line {i}: {line.strip()}")

# Appending to file
with open("test.txt", "a") as file:
    file.write("\nAppended line")

# JSON file handling
import json

data = {"name": "Alice", "age": 30, "skills": ["Python", "Java"]}
with open("data.json", "w") as file:
    json.dump(data, file, indent=2)

with open("data.json", "r") as file:
    loaded_data = json.load(file)
    print(f"Loaded data: {loaded_data}")

# Cleanup
os.remove("test.txt")
os.remove("data.json")
```

----------

## HTTP Requests
Making HTTP requests with requests library
```python
import requests
import json

# GET request
response = requests.get("https://httpbin.org/json")
print(f"Status: {response.status_code}")
print(f"JSON: {response.json()}")

# POST request with JSON
data = {"name": "Alice", "age": 30}
headers = {"Content-Type": "application/json"}

response = requests.post(
    "https://httpbin.org/post",
    json=data,
    headers=headers
)

if response.status_code == 200:
    result = response.json()
    print(f"POST response: {result['json']}")

# GET with parameters
params = {"page": 1, "limit": 10}
response = requests.get("https://httpbin.org/get", params=params)
print(f"URL with params: {response.url}")

# Error handling
try:
    response = requests.get("https://httpbin.org/status/404", timeout=5)
    response.raise_for_status()
except requests.exceptions.RequestException as e:
    print(f"Request error: {e}")
```

----------

## Flask Web Server
Creating a simple web server with Flask
```python
from flask import Flask, jsonify, request

app = Flask(__name__)

users = ["Alice", "Bob", "Charlie"]

@app.route("/")
def home():
    return "Hello, Flask!"

@app.route("/users", methods=["GET"])
def get_users():
    return jsonify(users)

@app.route("/users", methods=["POST"])
def create_user():
    data = request.json
    name = data.get("name")
    if name:
        users.append(name)
        return jsonify({"message": f"User {name} added"}), 201
    return jsonify({"error": "Name required"}), 400

if __name__ == "__main__":
    app.run(debug=True, port=5000)
```

----------

## String Manipulation
Common string operations
```python
text = "  Hello, World!  "
name = "Alice"
age = 30

# Basic operations
print(f"Original: '{text}'")
print(f"Trimmed: '{text.strip()}'")
print(f"Upper: {text.upper()}")
print(f"Lower: {text.lower()}")
print(f"Replace: {text.replace('World', 'Python')}")

# String formatting
formatted1 = "Name: {}, Age: {}".format(name, age)
formatted2 = f"Name: {name}, Age: {age}"
formatted3 = "Name: %(name)s, Age: %(age)d" % {"name": name, "age": age}

print(f"Formatted: {formatted2}")

# String methods
sentence = "python is awesome"
words = sentence.split()
capitalized = sentence.title()
starts_with = sentence.startswith("python")

print(f"Words: {words}")
print(f"Capitalized: {capitalized}")
print(f"Starts with 'python': {starts_with}")

# Join strings
joined = " | ".join(words)
print(f"Joined: {joined}")
```

----------

## Regular Expressions
Pattern matching with re module
```python
import re

text = "Contact us at info@company.com or support@company.com. Phone: (555) 123-4567"

# Email pattern
email_pattern = r"[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}"
emails = re.findall(email_pattern, text)
print(f"Found emails: {emails}")

# Phone pattern
phone_pattern = r"\((\d{3})\)\s(\d{3})-(\d{4})"
phone_match = re.search(phone_pattern, text)
if phone_match:
    area_code, prefix, number = phone_match.groups()
    print(f"Phone: ({area_code}) {prefix}-{number}")

# Replace pattern
masked_text = re.sub(email_pattern, "[EMAIL]", text)
print(f"Masked: {masked_text}")

# Split by pattern
parts = re.split(r"[.!?]", "Hello. How are you? I'm fine!")
print(f"Sentences: {[p.strip() for p in parts if p.strip()]}")

# Compile pattern for reuse
email_regex = re.compile(email_pattern)
matches = email_regex.finditer(text)
for match in matches:
    print(f"Email found at position {match.start()}-{match.end()}: {match.group()}")
```

----------

## Date and Time
Working with datetime module
```python
from datetime import datetime, date, time, timedelta
import time as time_module

# Current date and time
now = datetime.now()
today = date.today()
current_time = datetime.now().time()

print(f"Now: {now}")
print(f"Today: {today}")
print(f"Time: {current_time}")

# Formatting
formatted = now.strftime("%Y-%m-%d %H:%M:%S")
formatted_date = today.strftime("%B %d, %Y")

print(f"Formatted datetime: {formatted}")
print(f"Formatted date: {formatted_date}")

# Date arithmetic
tomorrow = today + timedelta(days=1)
next_week = now + timedelta(weeks=1)
one_hour_ago = now - timedelta(hours=1)

print(f"Tomorrow: {tomorrow}")
print(f"One hour ago: {one_hour_ago}")

# Parse date string
date_string = "2023-12-25 10:30:00"
parsed_date = datetime.strptime(date_string, "%Y-%m-%d %H:%M:%S")
print(f"Parsed: {parsed_date}")

# Timestamp
timestamp = now.timestamp()
from_timestamp = datetime.fromtimestamp(timestamp)
print(f"Timestamp: {timestamp}")
print(f"From timestamp: {from_timestamp}")
```

----------

## Collections and Data Structures
Advanced data structures
```python
from collections import defaultdict, Counter, deque, namedtuple
import heapq

# defaultdict
dd = defaultdict(list)
dd['fruits'].append('apple')
dd['fruits'].append('banana')
print(f"Default dict: {dict(dd)}")

# Counter
text = "hello world"
char_count = Counter(text)
print(f"Character count: {char_count}")
print(f"Most common: {char_count.most_common(3)}")

# deque (double-ended queue)
queue = deque([1, 2, 3])
queue.appendleft(0)
queue.append(4)
left = queue.popleft()
right = queue.pop()
print(f"Queue: {queue}, popped: {left}, {right}")

# namedtuple
Person = namedtuple('Person', ['name', 'age', 'city'])
person = Person('Alice', 30, 'New York')
print(f"Person: {person.name}, {person.age}")

# heapq (priority queue)
heap = [3, 1, 4, 1, 5, 9, 2, 6]
heapq.heapify(heap)
smallest = heapq.heappop(heap)
heapq.heappush(heap, 0)
print(f"Smallest: {smallest}, heap: {heap}")
```

----------

## Decorators
Function and class decorators
```python
import functools
import time

def timer(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        start = time.time()
        result = func(*args, **kwargs)
        end = time.time()
        print(f"{func.__name__} took {end - start:.4f} seconds")
        return result
    return wrapper

def retry(max_attempts=3):
    def decorator(func):
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            for attempt in range(max_attempts):
                try:
                    return func(*args, **kwargs)
                except Exception as e:
                    if attempt == max_attempts - 1:
                        raise e
                    print(f"Attempt {attempt + 1} failed: {e}")
        return wrapper
    return decorator

# Usage
@timer
@retry(max_attempts=2)
def unstable_function(x):
    if x < 0.5:
        raise ValueError("Random failure")
    return x * 2

# Property decorator
class Circle:
    def __init__(self, radius):
        self._radius = radius

    @property
    def radius(self):
        return self._radius

    @radius.setter
    def radius(self, value):
        if value < 0:
            raise ValueError("Radius cannot be negative")
        self._radius = value

    @property
    def area(self):
        return 3.14159 * self._radius ** 2

circle = Circle(5)
print(f"Area: {circle.area}")
```

----------

## Context Managers
Using and creating context managers
```python
import contextlib
import os

# File context manager (built-in)
with open("temp.txt", "w") as file:
    file.write("Hello, context manager!")

# Custom context manager class
class DatabaseConnection:
    def __init__(self, db_name):
        self.db_name = db_name
        self.connection = None

    def __enter__(self):
        print(f"Connecting to {self.db_name}")
        self.connection = f"connection_to_{self.db_name}"
        return self.connection

    def __exit__(self, exc_type, exc_val, exc_tb):
        print(f"Closing connection to {self.db_name}")
        if exc_type:
            print(f"Exception occurred: {exc_val}")
        return False

# Context manager using contextlib
@contextlib.contextmanager
def temporary_file(filename):
    print(f"Creating {filename}")
    try:
        with open(filename, "w") as f:
            yield f
    finally:
        print(f"Cleaning up {filename}")
        if os.path.exists(filename):
            os.remove(filename)

# Usage
with DatabaseConnection("users_db") as conn:
    print(f"Using {conn}")

with temporary_file("temp_data.txt") as f:
    f.write("Temporary data")
    print("File created and used")
```

----------

## Generators and Iterators
Creating generators and custom iterators
```python
# Generator function
def fibonacci(n):
    a, b = 0, 1
    for _ in range(n):
        yield a
        a, b = b, a + b

# Generator expression
squares = (x**2 for x in range(10))

# Custom iterator class
class CountDown:
    def __init__(self, start):
        self.start = start

    def __iter__(self):
        return self

    def __next__(self):
        if self.start <= 0:
            raise StopIteration
        self.start -= 1
        return self.start + 1

# Usage
print("Fibonacci sequence:")
for num in fibonacci(10):
    print(num, end=" ")

print("\nSquares:")
for square in squares:
    if square > 50:
        break
    print(square, end=" ")

print("\nCountdown:")
for num in CountDown(5):
    print(num, end=" ")

# Infinite generator
def infinite_sequence():
    num = 0
    while True:
        yield num
        num += 1

# Take first 5 from infinite sequence
gen = infinite_sequence()
first_five = [next(gen) for _ in range(5)]
print(f"\nFirst five: {first_five}")
```

----------

## Async Programming
Asynchronous programming with asyncio
```python
import asyncio
import time

async def say_hello(name, delay):
    await asyncio.sleep(delay)
    return f"Hello, {name}!"

async def fetch_data(data_id, delay):
    print(f"Fetching data {data_id}...")
    await asyncio.sleep(delay)
    return f"Data {data_id} fetched"

async def main():
    # Run tasks concurrently
    greet_tasks = [
        say_hello("Alice", 1),
        say_hello("Bob", 2),
        say_hello("Charlie", 1)
    ]
    
    # Sequential vs concurrent execution
    start = time.time()
    results = await asyncio.gather(*greet_tasks)
    concurrent_time = time.time() - start
    
    for result in results:
        print(result)
    
    print(f"Concurrent execution: {concurrent_time:.2f}s")
    
    # Fetch data concurrently
    data_tasks = [fetch_data(i, 0.5) for i in range(3)]
    data_results = await asyncio.gather(*data_tasks)
    print("Data results:", data_results)

if __name__ == "__main__":
    asyncio.run(main())
```

----------

## Testing with unittest
Unit testing framework
```python
import unittest
from unittest.mock import patch, MagicMock

def add(a, b):
    return a + b

def divide(a, b):
    if b == 0:
        raise ValueError("Cannot divide by zero")
    return a / b

class Calculator:
    def multiply(self, a, b):
        return a * b

class TestMathOperations(unittest.TestCase):
    def setUp(self):
        self.calculator = Calculator()

    def test_add(self):
        self.assertEqual(add(2, 3), 5)
        self.assertEqual(add(-1, 1), 0)

    def test_divide(self):
        self.assertEqual(divide(10, 2), 5.0)
        with self.assertRaises(ValueError):
            divide(10, 0)

    def test_multiply(self):
        self.assertEqual(self.calculator.multiply(3, 4), 12)

    @patch('requests.get')
    def test_mocked_request(self, mock_get):
        mock_response = MagicMock()
        mock_response.json.return_value = {"status": "ok"}
        mock_get.return_value = mock_response
        
        # Simulate function that uses requests
        import requests
        response = requests.get("https://api.example.com")
        self.assertEqual(response.json()["status"], "ok")

if __name__ == "__main__":
    unittest.main()
```

----------

## Command Line Arguments
Parsing command line arguments with argparse
```python
import argparse

def main():
    parser = argparse.ArgumentParser(description="Sample CLI app")
    
    parser.add_argument(
        "-n", "--name",
        default="World",
        help="Name to greet"
    )
    
    parser.add_argument(
        "-a", "--age",
        type=int,
        help="Age of the person"
    )
    
    parser.add_argument(
        "-v", "--verbose",
        action="store_true",
        help="Enable verbose output"
    )
    
    parser.add_argument(
        "files",
        nargs="*",
        help="Files to process"
    )
    
    args = parser.parse_args()
    
    if args.verbose:
        print(f"Arguments: {vars(args)}")
    
    greeting = f"Hello, {args.name}!"
    if args.age:
        greeting += f" You are {args.age} years old."
    
    print(greeting)
    
    if args.files:
        print(f"Processing files: {args.files}")

if __name__ == "__main__":
    main()
```

----------

## Logging
Structured logging with logging module
```python
import logging

# Basic logging setup
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s',
    handlers=[
        logging.FileHandler('app.log'),
        logging.StreamHandler()
    ]
)

logger = logging.getLogger(__name__)

def demo_logging():
    logger.debug("Debug message (won't show)")
    logger.info("Application started")
    logger.warning("This is a warning")
    logger.error("This is an error")
    
    try:
        result = 10 / 0
    except ZeroDivisionError:
        logger.exception("Division by zero occurred")

# Custom logger with different level
file_logger = logging.getLogger('file_logger')
file_logger.setLevel(logging.DEBUG)

file_handler = logging.FileHandler('debug.log')
file_handler.setFormatter(
    logging.Formatter('%(name)s - %(levelname)s - %(message)s')
)
file_logger.addHandler(file_handler)

# Usage
demo_logging()
file_logger.debug("This goes to debug.log")
logger.info("App running successfully")
```

----------

## Environment Variables
Working with environment variables and configuration
```python
import os

# Basic environment variable operations
debug_mode = os.getenv('DEBUG', 'False').lower() == 'true'
database_url = os.getenv('DATABASE_URL', 'sqlite:///app.db')
port = int(os.getenv('PORT', 5000))

# Required environment variable (raises KeyError if not set)
try:
    api_key = os.environ['API_KEY']
except KeyError:
    api_key = "default-api-key"

# Set environment variable
os.environ['TEMP_VAR'] = 'temporary_value'

# Check if variable exists
if 'HOME' in os.environ:
    print(f"Home: {os.environ['HOME']}")

# Configuration class
class Config:
    DEBUG = debug_mode
    DATABASE_URL = database_url
    PORT = port
    API_KEY = api_key

print(f"Debug mode: {Config.DEBUG}")
print(f"Port: {Config.PORT}")
print(f"Total env vars: {len(os.environ)}")
```

----------

## Data Processing with Pandas
Basic data manipulation with pandas
```python
import pandas as pd

# Create sample data
data = {
    'name': ['Alice', 'Bob', 'Charlie', 'Diana'],
    'age': [25, 30, 35, 28],
    'salary': [50000, 60000, 70000, 55000]
}

df = pd.DataFrame(data)
print(df)

# Filtering
young = df[df['age'] < 30]
high_earners = df[df['salary'] > 55000]

# Basic stats
avg_age = df['age'].mean()
total_salary = df['salary'].sum()

print(f"Average age: {avg_age}")
print(f"Total salary: {total_salary}")

# Add new column
df['salary_k'] = df['salary'] / 1000
df['category'] = df['age'].apply(lambda x: 'Young' if x < 30 else 'Senior')

# Sort and display
df_sorted = df.sort_values('age')
print(f"\nSorted by age:\n{df_sorted[['name', 'age', 'category']]}")
```