package main

import "fmt"

func main() {
	// ====================================== //
	//    			SubSlicing	       		  //
	// ====================================== //
	numbers := make([]int, 5)

	fmt.Println("========== numbers 5 ===========")
	fmt.Println("Slice is:", numbers)

	numbers = make([]int, 10)
	fmt.Println("========== numbers 10 ===========")
	fmt.Println("Slice is:", numbers)

	// ====================================== //
	//    			SubSlicing	       		  //
	// ====================================== //
	numbers = []int{0, 1, 2, 3, 4, 5, 6}
	sub1 := numbers[1:3]
	sub2 := numbers[3:]
	sub3 := numbers[4:]

	fmt.Println("========== numbers and subs ==========")
	fmt.Println("numbers :", numbers)
	fmt.Println("sub1:", sub1)
	fmt.Println("sub2:", sub2)
	fmt.Println("sub3:", sub3)
	fmt.Println("======================================")

	numbers = append(numbers, 0, 1, 2, 3, 4, 5)

	// Create a new slice with the same length and double capacity
	numbersMod := make([]int, len(numbers), cap(numbers)*2)

	//  Correct: copy from numbers → numbersMod
	copy(numbersMod, numbers)

	fmt.Println("Numbers: ", numbers)
	fmt.Println("Modified Numbers: ", numbersMod)

	// Nil Slice
	var num []int // nil slice
	fmt.Println("Nil Slices :", num)
	if num == nil {
		fmt.Println("Num is A Nil Silce")
	}
}
