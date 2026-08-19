package main

import (
	"fmt"
	"os"
)

// api_testing_suite - API testing tool
func api_testing_suite(path string) {
	fmt.Println("========================================")
	fmt.Println("  API-Testing-Suite")
	fmt.Println("  API testing tool")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	api_testing_suite(path)
}
