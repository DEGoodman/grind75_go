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

type ProblemTemplate struct {
	Number string
	Name   string
	Slug   string
}

func main() {
	// Only print debug message if G75_DEBUG is set
	if os.Getenv("G75_DEBUG") != "" {
		fmt.Println("DEBUG: Running grind75 CLI tool")
	}

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

	// Create template data
	tmplData := map[string]string{
		"__NUMBER__": number,
		"__NAME__":   name,
		"__SLUG__":   strings.ToLower(strings.ReplaceAll(name, " ", "-")),
	}

	// Process each template file
	templateFiles := []string{"solution.go", "solution_test.go", "README.md"}
	for _, filename := range templateFiles {
		// Read template
		templatePath := filepath.Join(templateDir, filename)
		tmplContent, err := os.ReadFile(templatePath)
		if err != nil {
			fmt.Printf("Error reading template %s: %v\n", filename, err)
			os.Exit(1)
		}

		// Replace template placeholders
		content := string(tmplContent)
		for placeholder, value := range tmplData {
			content = strings.ReplaceAll(content, placeholder, value)
		}

		// Write the processed template
		if err := os.WriteFile(filepath.Join(dirName, filename), []byte(content), 0644); err != nil {
			fmt.Printf("Error creating %s: %v\n", filename, err)
			os.Exit(1)
		}
	}

	fmt.Printf("Created new problem directory: %s\n", dirName)
}

func runTests(problemNumber string) {
	if os.Getenv("G75_DEBUG") != "" {
		fmt.Println("DEBUG: Starting runTests function")
	}

	if problemNumber == "all" {
		if os.Getenv("G75_DEBUG") != "" {
			fmt.Println("DEBUG: Running all tests")
		}
		cmd := exec.Command("go", "test", "./...")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running tests: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Find the problem directory using glob
	pattern := fmt.Sprintf("%s/%s_*", problemsDir, problemNumber)
	if os.Getenv("G75_DEBUG") != "" {
		fmt.Printf("DEBUG: Looking for tests matching pattern: %s\n", pattern)
	}

	matches, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("Error finding problem directory: %v\n", err)
		os.Exit(1)
	}
	if len(matches) == 0 {
		fmt.Printf("No problem found with number %s\n", problemNumber)
		os.Exit(1)
	}

	// Use the first match (there should only be one)
	testPath := fmt.Sprintf("./%s", matches[0])
	if os.Getenv("G75_DEBUG") != "" {
		fmt.Printf("DEBUG: Running tests in: %s\n", testPath)
	}

	cmd := exec.Command("go", "test", testPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running tests: %v\n", err)
		os.Exit(1)
	}
}

// findModuleRoot finds the directory containing go.mod by walking up from the current directory
func findModuleRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		// Check if go.mod exists in current directory
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// We've reached the root of the filesystem
			return "", fmt.Errorf("go.mod not found in any parent directory")
		}
		dir = parent
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

	fmt.Printf("Solution file: %s/solution.go\n", matches[0])
	fmt.Println("Note: Solutions are designed to be run through tests. Use 'grind75 test' to run the solution.")
}
