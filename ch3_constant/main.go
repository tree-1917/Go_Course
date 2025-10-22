package main

import (
	"fmt"
	"strings"
)

func main() {
	// === Numbers Example `Pizza Size` === //
	var block string = strings.Repeat("=", 20)
	const PI float32 = 3.243
	name := ""
	size := 0
	println(block)
	fmt.Println("=> Welcome Pizza System =>")
	println(block)

	fmt.Print("Enter Your Name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter Pizza Size: ")
	fmt.Scanf("%d", &size)

	// == Output == //
	price := float32(size) * PI

	fmt.Printf("Your Choice Pizza From size %d and Your Price $%0.2f\n", size, price)

	const (
		dec = 255
		oct = 03234
		hex = 0x324
	)
	println("Dec:", dec, "Oct:", oct, "Hex:", hex)
	const QUOTE = "\"GO is Simple!\" - A Programmer Can Learn It In One Day!"
	println(QUOTE)

	const MULTLINE = "Hello From Azzrk Company " +
		" It's Simple Software House" +
		" It's worker as salla partner"
	const ACTIVE = true
	const READY = false

	println("I am Active:", ACTIVE, "But I am Ready:", READY)
	println(MULTLINE)
	const LENGTH = 50
	const WIDTH = 10
	const AREA = LENGTH * WIDTH
	println("Area: ", AREA)
	println(block)
}
