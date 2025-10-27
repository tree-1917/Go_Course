package main

import (
	"errors" // Imported To Create Error Message
	"fmt"
	"math"
)

type error interface {
	Error() string
}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Negative nubmer passed to Sqrt")
	}
	return math.Sqrt(num), nil
}
func main() {
	// ======================== //
	// 		Error Handling 		//
	// ======================== //
	res, err := sqrt(-1)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", res)
	}

	res, err = sqrt(10)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", res)
	}
}
