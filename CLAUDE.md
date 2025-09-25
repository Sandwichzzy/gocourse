# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go learning repository (`gocourse`) containing example code demonstrating various Go language concepts and features. The codebase is organized into two main packages:

- **basics/**: Fundamental Go concepts (variables, functions, control flow, data structures)
- **intermediate/**: Advanced Go features (interfaces, generics, error handling, methods)

Each `.go` file in these directories contains standalone examples with their own `main()` function, designed for individual execution and learning.

## Architecture

### Package Structure
- `basics` package: Contains foundational Go programming examples
- `intermediate` package: Contains more advanced Go programming concepts
- Both packages are part of the `gocourse` module (Go 1.25.1)

### Key Concepts Covered
- **Basics**: Variables, functions, arrays, slices, maps, control flow, defer/panic/recover
- **Intermediate**: Interfaces, generics, custom errors, pointers, struct embedding, string manipulation

## Development Commands

### Running Individual Files
Since each `.go` file contains its own `main()` function, run them individually:
```bash
go run basics/functions.go
go run intermediate/interface.go
go run intermediate/generics.go
```

### Build Commands
- Cannot build the entire module due to multiple `main()` functions in each package
- Build individual files only: `go run <file>.go`

### Testing
- No test files exist in this codebase
- This is a learning repository focused on examples rather than production code

## Important Notes

- **Multiple main() functions**: Each package contains multiple files with `main()` functions, preventing package-level builds
- **Learning-focused**: Code is designed for educational purposes with examples and commented concepts
- **No external dependencies**: Uses only Go standard library
- **Individual execution**: Each file is meant to be run standalone to demonstrate specific concepts

## File Organization

Files are named descriptively based on the concept they demonstrate (e.g., `generics.go`, `interface.go`, `string_formatting.go`). Many files contain Chinese comments alongside English, indicating bilingual learning materials.