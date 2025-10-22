package main

import (
	"fmt"
	"strings"
)

func main() {
	block := strings.Repeat("*", 20)
	println(block)
	// ====================
	// ARITHMETIC OPERATORS
	// ====================
	A := 10
	B := 20

	fmt.Println("A + B =", A+B)
	fmt.Println("A - B =", A-B)
	fmt.Println("A * B =", A*B)
	fmt.Println("A / B =", A/B)
	fmt.Println("A % B =", A%B)
	A++
	fmt.Println("A++ =", A)
	A--
	fmt.Println("A-- =", A)

	println(block)
	// ====================
	// Relation OPERATORS
	// ====================
	fmt.Println("A == B ~>", A == B)
	fmt.Println("A != B ~>", A != B)
	fmt.Println("A > B ~>", A > B)
	fmt.Println("A < B ~>", A < B)
	fmt.Println("A >= B ~>", A >= B)
	fmt.Println("A <= B ~>", A <= B)

	println(block)
	// ====================
	// Logical OPERATORS
	// ====================
	AA := true
	BB := false
	fmt.Println("AA && BB =>", AA && BB)
	fmt.Println("AA || BB =>", AA || BB)
	fmt.Println("!AA  =>", !AA)
	fmt.Println("!BB =>", !BB)

	println(block)
	// ====================
	// Assignment OPERATORS
	// ====================
	A += B
	fmt.Println("A += B =>", A)

	A *= B
	fmt.Println("A *= B =>", A)

	A -= B
	fmt.Println("A -= B =>", A)

	println(block)
	// ====================
	// Bitwise OPERATORS
	// ====================
	A &= 2 // A = A & 2 --> 1010
	// 				0010
	//				----
	// 				0010 => 2
	println("A &=2:", A)

	A |= 2 // A = A | 2 -> 1010
	// 			  0010
	// 			  ----
	// 			  1010

	println("A |=2:", A)

	println(block)
	// =======================
	// Miscellaneous OPERATORS
	// =======================
	num := 10
	ptr := &num

	fmt.Println("Address of A:", ptr)

	fmt.Println("Value of *ptr:", *ptr)
	println(block)
}
