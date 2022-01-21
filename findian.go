package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userInput string
	fmt.Println("Enter a string: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
	}

	var result string
	if containsIan(userInput) {
		result = "Found!"
	} else {
		result = "Not Found!"
	}

	fmt.Printf("%s", result)
}

func containsIan(text string) bool {
	return (text[0] == 'i' || text[0] == 'I') && (text[len(text)-1] == 'n' || text[len(text)-1] == 'N') && (strings.Contains(text, "a") || strings.Contains(text, "A"))
}
