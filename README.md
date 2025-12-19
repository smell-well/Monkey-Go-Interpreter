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

This implementation currently includes:

- **Lexer**: Tokenizes Monkey source code into a stream of tokens
- **Parser**: Builds an Abstract Syntax Tree (AST) from tokens using Pratt parsing
  - Let statements (`let x = 5;`)
  - Return statements (`return x;`)
  - Expression statements
  - Integer literals
  - Prefix expressions (`-5`, `!true`)
  - Infix expressions (`5 + 5`, `x * y`)
- **AST**: Complete node structure for representing Monkey programs
- **REPL**: Read-Eval-Print-Loop for interactive experimentation

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
>> -50 + 100
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests for specific packages
go test ./lexer
go test ./parser
go test ./ast
```

## Project Structure

```
.
├── main.go          # Entry point, starts the REPL
├── token/           # Token definitions and types
├── lexer/           # Lexical analyzer
├── ast/             # Abstract Syntax Tree node definitions
├── parse/           # Parser implementation
└── repl/            # Read-Eval-Print-Loop
```

## Example Monkey Code

```javascript
// Variable bindings
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);

// Expressions
!true;
5 + 5 * 10;
-5;

// Return statements
return 5;
return 10 + 15;
```

## Development Progress

- [x] Lexer
- [x] Parser (Pratt parsing)
- [x] AST representation
- [x] REPL
- [ ] Evaluator
- [ ] Object system
- [ ] Environment and bindings
- [ ] Functions
- [ ] Built-in functions
- [ ] Arrays
- [ ] Hashes

## Learning Resources

This project is based on [Writing An Interpreter In Go](https://interpreterbook.com/) by Thorsten Ball. The book provides a comprehensive guide to building an interpreter from scratch.

## License

This is a learning project. Feel free to use it for educational purposes.

## Acknowledgments

- Thorsten Ball for the excellent book _Writing An Interpreter In Go_
- The Go community for providing great tools and documentation
