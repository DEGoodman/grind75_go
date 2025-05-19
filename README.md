# Grind75 Go Solutions

This repository contains my solutions to the Grind75 problem set implemented in Go. The project includes a simple CLI tool to help manage and test solutions efficiently.

## Project Structure

```
.
├── cmd/
│   └── grind75/         # CLI tool for managing Grind75 problems
├── problems/            # Directory containing all Grind75 solutions
│   └── template/        # Template for new problems
├── go.mod              # Go module file
└── README.md           # This file
```

## Setup

1. Make sure you have Go installed (version 1.16 or later recommended)
2. Clone this repository
3. Install the CLI tool:
   ```bash
   go install ./cmd/grind75
   ```

## Usage

### Creating a New Problem

To create a new problem solution:

```bash
grind75 new <problem_number> <problem_name>
```

For example:
```bash
grind75 new 1 "Two Sum"
```

This will create a new directory in `problems/` with:
- A solution file (`solution.go`)
- A test file (`solution_test.go`)
- A README.md with problem description template

### Running Tests

To run tests for a specific problem:

```bash
grind75 test <problem_number>
```

For example:
```bash
grind75 test 1
```

To run all tests:

```bash
grind75 test all
```

### Running a Solution

To run a specific solution:

```bash
grind75 run <problem_number>
```

For example:
```bash
grind75 run 1
```

## Development

The project uses Go modules for dependency management. To add new dependencies:

```bash
go get <package_name>
```

## Contributing

Feel free to use this template for your own Grind75 practice. The CLI tool is designed to be simple and extensible. 