package main

import (
	"fmt"
	"strings"
)

func main() {
	// ========================================= //
	// ========== IF STATEMENT ================= //
	// ========================================= //

	block := strings.Repeat("█", 60)
	sblock := strings.Repeat("⭐️", 30)
	println(block)
	println(sblock)
	// var (
	// 	age   int
	// 	name  string
	// 	score int
	// )
	//
	// fmt.Print("Enter your name: ")
	// _, err := fmt.Scanln(&name)
	// if err != nil {
	// 	panic("Cannot take name as input")
	// }
	//
	// fmt.Print("Enter your age: ")
	// _, err = fmt.Scanf("%d", &age)
	// if err != nil {
	// 	panic("Cannot take age as input")
	// }
	// fmt.Print("Enter your score: ")
	// _, err = fmt.Scanf("%d", &score)
	// if err != nil {
	// 	panic("Cannot take score as input")
	// }

	// if age >= 18 && score >= 80 {
	// 	fmt.Printf("Hello %s, you are eligible to vote twice.\n", name)
	// } else if age >= 18 && score >= 50 {
	// 	fmt.Printf("Hello %s, you are eligible to vote once.\n", name)
	// } else {
	// 	fmt.Printf("Hello %s, you aren't eligible to vote.\n", name)
	// }
	// println(block)

	// ============================================ //
	// ========== Loops STATEMENT ================= //
	// ============================================ //

	// -----------
	// For Loop
	// -----------
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j < 10; j++ {
	// 		fmt.Printf("%d * %d = %d\t", i, j, i*j)
	// 	}
	// 	fmt.Println()
	// }

	// ------------------------------------------- //
	// For Loop (break) (continue) (infinite loop) //
	// ------------------------------------------- //
	target := 13
	var name string
	print("Enter Your Name: ")

	fmt.Scanln(&name)

	var (
		userVal int
		state   bool = false
		tries   int  = 3
	)
	println(sblock)
	println("Guess Number From 1 to 20: ")
	println(sblock)
	for {

		if tries == 0 {
			break
		}

		print("Enter Your Guess: ")
		fmt.Scanf("%d", &userVal)

		if userVal == 77 {
			continue
		} else if userVal == 97 {
			goto backdoor
		} else {
			tries--
		}

		if userVal == target {
			state = true
		} else if userVal > target {
			println("So Big")
		} else if userVal < target {
			println("So Small")
		}

	}
	println(sblock)
	if state {
		println("Success")
	} else {
		println("Failed")
	}
backdoor:
	println("Success")
}
