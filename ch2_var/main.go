package main

import (
	"fmt"
	"strings"
)

func main() {
	block := strings.Repeat("=", 20)
	hblock := strings.Repeat("*", 50)
	fmt.Println(hblock)
	// ================================================
	// way1: Declare and assign on 2 different lines
	// ================================================
	var mango string
	mango = "This is a mango!"
	var weight int
	weight = 32

	fmt.Println(block, "Way 1", block)
	fmt.Println("mango: ", mango)
	fmt.Println("weight:", weight)

	// ================================================
	// way2: Declare and assign on same line
	// ================================================
	var height int = 23
	var applesNum, orangesNum int = 32, 123
	var a, b, c = 123.12, 32, "user1"

	fmt.Println("\n"+block, "Way 2", block)
	fmt.Println("mango: ", mango)
	fmt.Println("weight:", weight)
	fmt.Println("height:", height)
	fmt.Println("Apple Number:", applesNum)
	fmt.Println("Orange Number:", orangesNum)
	fmt.Printf("a:%f\tb:%d\tc:%s", a, b, c)

	// ================================================
	// way3: shorthand
	// ================================================
	age := 32
	applesNum1, orangesNum1 := 12, 13
	win, mac, linux := "win is no", "mac is no no", "linux is go"
	fruitsNum := applesNum1 + orangesNum1

	fmt.Println("\n"+block, "Way 3", block)
	fmt.Println("mango: ", mango)
	fmt.Println("weight:", weight)
	fmt.Println("height:", height)
	fmt.Println("Apple Number:", applesNum)
	fmt.Println("Orange Number:", orangesNum)
	fmt.Println("Apple Number(go):", applesNum1)
	fmt.Println("Orange Number(go):", orangesNum1)
	fmt.Println("fruitsNum (go):", fruitsNum)
	fmt.Println(block)
	fmt.Printf("Current News: \n%s\n%s\n%s\n", win, mac, linux)
	fmt.Println(block)
	fmt.Println("Age :", age)
	fmt.Printf("Age is of type: %T\n", age)
	fmt.Println(hblock)

}
