package main

import (
	"fmt"
	"strings"
)

// Pass By Pointer
func swap(fname, lname *string) { *fname, *lname = *lname, *fname }

func main() {
	fullName := "Ali Hassan"
	parts := strings.Fields(fullName)

	var firstName, lastName string

	if len(parts) >= 2 {
		firstName = parts[0]
		lastName = parts[1]
	} else {
		fmt.Println("Invalid full name")
		return
	}
	fmt.Println("==================================")
	fmt.Println("=========== Before Swap ==========")
	fmt.Println("==================================")
	fmt.Printf("=> %s %s\n", firstName, lastName)
	fmt.Println("==================================")
	fmt.Println("=========== Before Swap ==========")
	fmt.Println("==================================")
	swap(&firstName, &lastName)
	fmt.Printf("=> %s %s\n", firstName, lastName)
}
