package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	templateDir = "problems/template"
	problemsDir = "problems"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "new":
		if len(os.Args) < 4 {
			fmt.Println("Error: 'new' command requires problem number and name")
			printUsage()
			os.Exit(1)
		}
		createNewProblem(os.Args[2], os.Args[3])
	case "test":
		if len(os.Args) < 3 {
			fmt.Println("Error: 'test' command requires problem number")
			printUsage()
			os.Exit(1)
		}
		runTests(os.Args[2])
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Error: 'run' command requires problem number")
			printUsage()
			os.Exit(1)
		}
		runSolution(os.Args[2])
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  grind75 new <problem_number> <problem_name>  - Create a new problem solution")
	fmt.Println("  grind75 test <problem_number>               - Run tests for a problem")
	fmt.Println("  grind75 test all                            - Run all tests")
	fmt.Println("  grind75 run <problem_number>                - Run a solution")
}

func createNewProblem(number, name string) {
	// Create problem directory
	dirName := fmt.Sprintf("%s/%s_%s", problemsDir, number, strings.ToLower(strings.ReplaceAll(name, " ", "_")))
	if err := os.MkdirAll(dirName, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Create solution.go
	solutionContent := fmt.Sprintf(`package main

// Problem %s: %s
// https://leetcode.com/problems/%s/

func main() {
	// TODO: Implement your solution here
}
`, number, name, strings.ToLower(strings.ReplaceAll(name, " ", "-")))

	if err := os.WriteFile(filepath.Join(dirName, "solution.go"), []byte(solutionContent), 0644); err != nil {
		fmt.Printf("Error creating solution.go: %v\n", err)
		os.Exit(1)
	}

	// Create solution_test.go
	testContent := fmt.Sprintf(`package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		// TODO: Add your test cases here
		{
			name:     "Example 1",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: Implement your test cases
			t.Skip("Test not implemented yet")
		})
	}
}
`)

	if err := os.WriteFile(filepath.Join(dirName, "solution_test.go"), []byte(testContent), 0644); err != nil {
		fmt.Printf("Error creating solution_test.go: %v\n", err)
		os.Exit(1)
	}

	// Create README.md
	readmeContent := fmt.Sprintf(`# Problem %s: %s

## Problem Description
[Add problem description here]

## Solution
[Add solution explanation here]

## Time Complexity
[Add time complexity analysis here]

## Space Complexity
[Add space complexity analysis here]
`, number, name)

	if err := os.WriteFile(filepath.Join(dirName, "README.md"), []byte(readmeContent), 0644); err != nil {
		fmt.Printf("Error creating README.md: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created new problem directory: %s\n", dirName)
}

func runTests(problemNumber string) {
	if problemNumber == "all" {
		cmd := exec.Command("go", "test", "./...")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running tests: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Find the problem directory
	matches, err := filepath.Glob(fmt.Sprintf("%s/%s_*", problemsDir, problemNumber))
	if err != nil {
		fmt.Printf("Error finding problem directory: %v\n", err)
		os.Exit(1)
	}
	if len(matches) == 0 {
		fmt.Printf("No problem found with number %s\n", problemNumber)
		os.Exit(1)
	}

	cmd := exec.Command("go", "test", matches[0])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running tests: %v\n", err)
		os.Exit(1)
	}
}

func runSolution(problemNumber string) {
	// Find the problem directory
	matches, err := filepath.Glob(fmt.Sprintf("%s/%s_*", problemsDir, problemNumber))
	if err != nil {
		fmt.Printf("Error finding problem directory: %v\n", err)
		os.Exit(1)
	}
	if len(matches) == 0 {
		fmt.Printf("No problem found with number %s\n", problemNumber)
		os.Exit(1)
	}

	cmd := exec.Command("go", "run", filepath.Join(matches[0], "solution.go"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running solution: %v\n", err)
		os.Exit(1)
	}
} 