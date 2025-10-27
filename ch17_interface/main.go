package main

import "fmt"

/*
	- Interface :
		- An interface type is defined as a "set of method signatures" .
		- A Value of interface type can hold any value that implements those methods .
*/

// Define an interface
type Shape interface {
	area() float64
}

func getArea(s Shape) float64 {
	return s.area()
}

// ========================= //
//
//	Rectangle Shape  	 //
//
// ========================= //
// Define an Rectangle struct
type Rect struct {
	width, height float64
}

// Implment Area Method For Circle
func (r Rect) area() float64 {
	return r.height * r.width
}

// ========================= //
//
//	Circle Shape 		 //
//
// ========================= //
// Define an Circle Struct
type Circle struct {
	radius float64
}

// Implment Area Method For Circle
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	// ==================== //
	// 		Interface 		//
	// ==================== //

	myCircle := Circle{radius: 5}
	myRect := Rect{width: 10, height: 10}

	fmt.Printf("Circle area %.02f\n", getArea(myCircle))
	fmt.Printf("Rectangle area %.02f\n", getArea(myRect))
}
