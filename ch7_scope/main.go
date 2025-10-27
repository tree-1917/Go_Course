package main

import "fmt"

// 2. Global Vars ( All Functions in Same Package See It )
var g int
var x int

func main() {
	// 1. local Variables ( Only Functions In It's Scoope See It)
	var a, b, c int

	a = 10
	b = 20
	c = a + b
	g = a * b
	var x int = 100
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
	fmt.Printf("a = %d, b = %d, g = %d\n", a, b, g)
	fmt.Printf("x = %d\n", x)
}
