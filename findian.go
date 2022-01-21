package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput string
	fmt.Println("Enter a string: ")

	fmt.Scan(&userInput)

	var result string
	if containsIan(userInput) {
		result = "Found!"
	} else {
		result = "Not Found!"
	}

	fmt.Printf("%s", result)
}

func containsIan(text string) bool {
	return text[0] == 'i' && text[len(text)-1] == 'n' && strings.Contains(text, "a")
}
