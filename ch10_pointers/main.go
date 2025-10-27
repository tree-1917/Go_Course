package main

import "fmt"

func main() {
	// Understanding Pointers
	// A pointer stores the memory address of another variable .
	// Declare a Pointer
	var ip *int
	var fp *float32

	// Basic Pointer Operation
	// 1. Store Address & address of
	var i int = 10
	var ii float32 = 30.204

	ip = &i
	fp = &ii

	fmt.Printf("[%v] => %d\n", ip, i)
	fmt.Printf("[%v] => %f\n", fp, ii)

	var nums = []int{1, 2, 3}

	var ptr *int = &nums[0]
	fmt.Println("==========")
	fmt.Printf("[%v]=> %d\n", ptr, nums[0])
}
