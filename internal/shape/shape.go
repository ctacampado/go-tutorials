package shape

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Square struct where X one side of the square
type Square struct {
	X float64
}

// Area returns the area of the square
func (s Square) Area() float64 {
	return s.X * s.X
}

// Perimeter returns the perimeter of the square
func (s Square) Perimeter() float64 {
	return 4 * s.X
}

// Circle struct where R is the radius
type Circle struct {
	R float64
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return c.R * c.R * math.Pi
}

// Perimeter returns the perimeter of the square
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

// Calculate identifies the argument shape and
// calculates its area and perimeter
func Calculate(s Shape) {
	_, ok := s.(Circle) // type assertion
	if ok {
		fmt.Println("Is a circle!")
	}

	v, ok := s.(Square) // type assertion
	if ok {
		fmt.Println("Is a square!", v)
	}

	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
}
