package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString string = "Hello, World!"

	for i, c := range myString {
		fmt.Printf("Index: %d, Character: %c\n", i, c)
	}

	var stringSlice []string = []string{"h", "e", "l", "l", "o"}
	var catString string = strings.Join(stringSlice, "")

	fmt.Println(catString)
}
