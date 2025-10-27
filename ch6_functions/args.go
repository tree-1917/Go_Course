package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main <name> <age>")
		return
	}
	name := os.Args[1]
	age := os.Args[2]

	fmt.Printf("Hello %s! Your are %s years old.\n", name, age)
}
