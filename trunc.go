package main

import "fmt"

func main() {
	var userFloat float32
	fmt.Println("Enter a floating point number: ")

	fmt.Scan(&userFloat)

	fmt.Printf("%.0f", userFloat)
}
