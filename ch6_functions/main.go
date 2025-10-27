package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func add(num1, num2 int) int { return num1 + num2 }
func sub(num1, num2 int) int { return num1 - num2 }
func mul(num1, num2 int) int { return num1 * num2 }
func div(num1, num2 int) int { return num1 / num2 }

func main() {
	fmt.Println("------------------------------")

	reader := bufio.NewReader(os.Stdin)

	var (
		num1, num2 int
		opt        rune
		islove     rune = 'y'
	)

	for {
		if unicode.ToLower(islove) == 'n' {
			break
		}

		fmt.Print("Enter Your First Num: ")
		fmt.Scan(&num1)

		fmt.Print("Ops (+,-,*,/): ")
		opt, _, _ = reader.ReadRune() // read operator cleanly
		reader.ReadString('\n')       // clear buffer

		fmt.Print("Enter Your Second Num: ")
		fmt.Scan(&num2)

		switch opt {
		case '+':
			fmt.Printf("%d + %d = %d\n", num1, num2, add(num1, num2))
		case '-':
			fmt.Printf("%d - %d = %d\n", num1, num2, sub(num1, num2))
		case '*':
			fmt.Printf("%d * %d = %d\n", num1, num2, mul(num1, num2))
		case '/':
			fmt.Printf("%d / %d = %d\n", num1, num2, div(num1, num2))
		default:
			fmt.Println("Error: invalid operator")
		}

		fmt.Print("Do you love (y/n): ")
		islove, _, _ = reader.ReadRune()
		reader.ReadString('\n') // clear buffer
		islove = unicode.ToLower(islove)
	}

	fmt.Println("------------------------------")
}
