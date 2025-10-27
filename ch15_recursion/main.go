package main

import "fmt"

func main() {
	// ===================== //
	// 		Recursion 		 //
	// ===================== //
	f10 := factorial(10)

	fmt.Println("Factorial of 10", f10)
}

// + Factorial {{{
func factorial(n int) int {
	if n == 0 {
		return 1 // Base Case
	} else {
		return n * factorial(n-1) // Recursion Case
	}
}

// }}}
