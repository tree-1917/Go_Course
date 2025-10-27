package main

import "fmt"

func main() {
	var total int = 30
	var items int = 7
	var avg float64

	avg = float64(total) / float64(items)

	fmt.Printf("Avg : %0.2f\n", avg)
}
