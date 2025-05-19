# Grind75 Go Solutions

This repository contains my solutions to the (Grind75 problem set)[https://www.techinterviewhandbook.org/grind75/] implemented in Go. The project includes a simple CLI tool to help manage and test solutions efficiently.

## Project Structure

```
.
├── cmd/
│   └── grind75/         # CLI tool for managing Grind75 problems
├── problems/            # Directory containing all Grind75 solutions
│   ├── template/        # Template files for new problems
│   │   ├── solution.go      # Solution template with proper function signature
│   │   ├── solution_test.go # Test template with table-driven tests
│   │   └── README.md        # Problem documentation template
│   └── [problem_dirs]/  # Individual problem solutions
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
- A solution file (`solution.go`) with:
  - Proper function signature and `return nil` stub
  - Example usage in `main()`
  - Problem documentation
- A test file (`solution_test.go`) with:
  - Table-driven test structure
  - Basic assertions for length and values
  - Placeholder for additional test cases
- A README.md with:
  - Problem description template
  - Sections for solution explanation and complexity analysis
  - Notes about error handling conventions

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

### Development Tips

1. **Debug Mode**
   Set the `G75_DEBUG` environment variable to see detailed debug output:
   ```bash
   export G75_DEBUG=1
   grind75 test 1
   ```

2. **Test Development**
   - Use `go test -v ./problems/1_two_sum` to run tests with verbose output
   - Use `go test -run TestTwoSum/Example_1 ./problems/1_two_sum` to run a specific test case
   - Add `-count=1` to disable test caching: `go test -count=1 ./problems/1_two_sum`

3. **Code Organization**
   - Each problem is in its own package (e.g., `problem_1`)
   - Test files are in the same package as the solution
   - Use table-driven tests for multiple test cases
   - Document time and space complexity in the README

4. **Error Handling**
   - See the "Error Handling Conventions" section below
   - Focus on the algorithmic challenge rather than input validation
   - Use clear error messages in tests

### Template System

The project includes a template system for creating new problem solutions. Templates are stored in `problems/template/` and include:

1. `solution.go`:
   - Proper function signature with `return nil` stub
   - Example usage in `main()`
   - Problem documentation with LeetCode link

2. `solution_test.go`:
   - Table-driven test structure
   - Typed input and expected values
   - Basic assertions for length and values
   - Placeholder for additional test cases

3. `README.md`:
   - Problem description template
   - Sections for solution explanation
   - Time and space complexity analysis
   - Notes about error handling conventions

The CLI tool uses Go's `text/template` package to process these templates, replacing placeholders like `{{.Number}}`, `{{.Name}}`, and `{{.Slug}}` with the appropriate values.

## Error Handling Conventions

This project intentionally deviates from Go's idiomatic error handling patterns to conform to LeetCode's problem format. Specifically:

- Solutions return only the expected result type (e.g., `[]int` for Two Sum) rather than the idiomatic Go pattern of `(result, error)`
- Input validation is assumed to be handled by LeetCode's test environment
- Error cases that would typically return errors in production code (like invalid inputs) are not explicitly handled

This design choice is made to:
1. Match LeetCode's expected function signatures
2. Keep solutions focused on the core algorithmic challenge
3. Maintain consistency with LeetCode's problem format

In production Go code, you would typically want to:
- Return `(result, error)` for functions that can fail
- Validate inputs explicitly
- Handle edge cases with proper error returns
- Use custom error types for different failure modes

E.G.: 
```
func doSomething() (Result, error) {
    // ... do work
    if someCondition {
        return Result{}, fmt.Errorf("something went wrong: %w", err)
    }
    return result, nil
}

func caller() error {
    result, err := doSomething()
    if err != nil {
        return fmt.Errorf("failed to do something: %w", err)
    }
    // use result
    return nil
}
```

## Contributing

Feel free to use this template for your own Grind75 practice. The CLI tool is designed to be simple and extensible. 