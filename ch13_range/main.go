package main

import "fmt"

func main() {
	// ================ //
	// 		Range 		//
	// ================ //

	numbers := []int{0, 1, 2, 3, 4, 5}

	for ix, number := range numbers {
		fmt.Printf("[%d] => %d \n", ix, number)
	}
}
