package main

import (
	"fmt"
	"strings"
)

func main() {
	// ====================
	// 		Strings
	// ====================
	var greetings string = "Hello, World!"

	fmt.Printf("normal string: %s\n", greetings)

	// Display The Hexdecimal byte values
	fmt.Printf("Hex Bytes: ")

	for i := 0; i < len(greetings); i++ {
		fmt.Printf("%x", greetings[i])
	}
	fmt.Println("")
	fmt.Println("the length of this string :", len(greetings))
	// String Concatentation
	fruits := []string{"apples", "orange", "bananas"}
	fmt.Println("I eat " + strings.Join(fruits, " "))
}
