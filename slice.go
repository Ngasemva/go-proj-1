package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := make([]int, 0, 3)

	for {
		var userInput string
		fmt.Println("Enter an integer to be added to the slice, or X/x to quit: ")

		fmt.Scan(&userInput)

		intToAdd, err := strconv.Atoi(userInput)

		if err != nil {
			if strings.ToLower(userInput) == "x" {
				fmt.Println("Stopping the program")
				break
			} else {
				fmt.Printf("The only acceptable inputs are integers or x/X, Your input %s was not accepted\n", userInput)
				continue
			}
			fmt.Println(userInput)
			break
		}
		s = append(s, intToAdd)

		sort.Ints(s)

		fmt.Printf("\n%v\n", s)
	}
}
