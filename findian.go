package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func old_findian() {
	var strInput string
	fmt.Println("Input the string : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strInput = scanner.Text()
	var lowerStrInput string = strings.ToLower(strInput)
	fmt.Println(lowerStrInput)
	if strings.HasPrefix(lowerStrInput, "i") && strings.HasSuffix(lowerStrInput, "n") && strings.Contains(lowerStrInput, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
