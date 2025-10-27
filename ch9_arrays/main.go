package main

import "fmt"

func main() {
	// Declareing Arrays
	var xbalance = []float32{12.4, 23.23, 49.34, 64.34, 10.20, 20.04}
	for i, val := range xbalance {
		fmt.Printf("Element[%d] = %f\n", i, val)
	}
}
