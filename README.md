# Monkey Programming Language

A Go implementation of the Monkey programming language interpreter, based on the book [_Writing An Interpreter In Go_](https://interpreterbook.com/) by Thorsten Ball.

## About Monkey

Monkey is a programming language designed for learning how to build an interpreter from scratch. It features:

- C-like syntax
- Variable bindings
- Integer and boolean data types
- Arithmetic expressions
- Built-in functions
- First-class and higher-order functions
- Closures
- String data structure
- Array data structure
- Hash data structure

## Current Implementation Status

This is a complete implementation of the Monkey interpreter with all features from the book:

### Core Components

- **Lexer**: Tokenizes Monkey source code into a stream of tokens
- **Parser**: Builds an Abstract Syntax Tree (AST) from tokens using Pratt parsing
- **Evaluator**: Tree-walking interpreter that evaluates the AST
- **Object System**: Runtime representation of values
- **Environment**: Variable binding and scope management
- **REPL**: Read-Eval-Print-Loop for interactive experimentation

### Supported Features

- **Data Types**: Integers, Booleans, Strings, Arrays, Hash Maps, Null
- **Operators**: Arithmetic (`+`, `-`, `*`, `/`), Comparison (`==`, `!=`, `<`, `>`), Logical (`!`)
- **Variable Bindings**: `let` statements
- **Functions**: First-class functions, closures, higher-order functions
- **Control Flow**: `if-else` expressions
- **Return Statements**: Early returns from functions
- **Built-in Functions**:
  - `len()`: Get length of strings or arrays
  - `first()`: Get first element of array
  - `last()`: Get last element of array
  - `rest()`: Get all elements except first
  - `push()`: Add element to array
  - `put()`: Print to console

## Installation

Make sure you have Go 1.24.5 or later installed.

```bash
# Clone the repository
git clone https://github.com/smell-well/Monkey-Go-Interpreter.git
cd Monkey-Go-Interpreter

# Run the REPL
go run main.go
```

## Usage

### Interactive REPL

Start the REPL to experiment with Monkey code:

```bash
go run main.go
```

Example session:

```
Hello your_name! This is the Monkey programming language!
Feel free to type in commands
>> let x = 5;
>> x + 10
15
>> let add = fn(a, b) { a + b };
>> add(5, 10)
15
>> let map = fn(arr, f) { let iter = fn(arr, accumulated) { if (len(arr) == 0) { accumulated } else { iter(rest(arr), push(accumulated, f(first(arr)))); } }; iter(arr, []); };
>> let double = fn(x) { x * 2 };
>> map([1, 2, 3], double)
[2, 4, 6]
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests for specific packages
go test ./lexer
go test ./parse
go test ./ast
go test ./evaluator
go test ./object
```

## Project Structure

```
.
├── main.go          # Entry point, starts the REPL
├── token/           # Token definitions and types
├── lexer/           # Lexical analyzer
├── ast/             # Abstract Syntax Tree node definitions
├── parse/           # Parser implementation (Pratt parser)
├── evaluator/       # Tree-walking interpreter
├── object/          # Object system and runtime values
│   ├── object.go    # Object types (Integer, String, Function, etc.)
│   └── environment.go # Variable binding and scope
└── repl/            # Read-Eval-Print-Loop
```

## Example Monkey Code

```javascript
// Variable bindings
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);

// Arithmetic and boolean expressions
5 + 5 * 10;
!true;
if (10 > 5) { true } else { false }

// Functions
let add = fn(a, b) { return a + b; };
add(5, 10);

// Higher-order functions
let applyFunc = fn(a, b, func) { func(a, b) };
applyFunc(2, 3, add);

// Closures
let newAdder = fn(x) { fn(y) { x + y } };
let addTwo = newAdder(2);
addTwo(3); // Returns 5

// Strings
let greeting = "Hello" + " " + "World!";
len(greeting);

// Arrays
let myArray = [1, 2, 3, 4];
first(myArray);
last(myArray);
rest(myArray);
push(myArray, 5);

// Hash maps
let person = {"name": "Alice", "age": 30};
person["name"];

// Recursive functions (Fibonacci)
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      1
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};
fibonacci(10);

// Map function
let map = fn(arr, f) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0) {
      accumulated
    } else {
      iter(rest(arr), push(accumulated, f(first(arr))));
    }
  };
  iter(arr, []);
};
let double = fn(x) { x * 2 };
map([1, 2, 3, 4], double); // Returns [2, 4, 6, 8]
```

## Development Progress

- [x] Lexer
- [x] Parser (Pratt parsing)
- [x] AST representation
- [x] REPL
- [x] Evaluator (Tree-walking interpreter)
- [x] Object system
- [x] Environment and bindings
- [x] Functions and closures
- [x] Built-in functions (len, first, last, rest, push, put)
- [x] Strings
- [x] Arrays
- [x] Hash maps
- [x] Error handling

**Status**: Complete implementation of "Writing An Interpreter In Go"

## Key Features & Highlights

- **Complete Implementation**: All features from the book are fully implemented
- **Well-Tested**: Comprehensive test suites for lexer, parser, evaluator, and object system
- **Clean Architecture**: Separation of concerns between lexing, parsing, and evaluation
- **Pratt Parsing**: Elegant operator precedence parsing
- **Tree-Walking Interpreter**: Direct AST evaluation without bytecode compilation
- **First-Class Functions**: Functions are values that can be passed around and returned
- **Closures**: Functions can capture and access variables from their surrounding scope
- **Immutable Data Structures**: Arrays and hashes create new copies on modification

## Learning Resources

This project is based on [Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball. The book provides a comprehensive guide to building an interpreter from scratch.

### What You'll Learn

- Lexical analysis and tokenization
- Parsing techniques (Pratt parsing/Top-Down Operator Precedence)
- Abstract Syntax Trees (AST)
- Tree-walking interpretation
- Object systems and runtime representation
- Environment and scope management
- First-class functions and closures
- Built-in function implementation

## License

This is a learning project. Feel free to use it for educational purposes.

## Acknowledgments

- Thorsten Ball for the excellent book _Writing An Interpreter In Go_
- The Go community for providing great tools and documentation
