package main

import (
	"fmt"
	"strconv"
)

func main() {
	var total int = 30
	var items int = 7
	var avg float64

	avg = float64(total) / float64(items)

	fmt.Printf("Avg : %0.2f\n", avg)

	// Converting String To Integer
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error Converting", err)
		return
	}
	fmt.Printf("Converted Number : %v\n", num)
}
