package main

import "fmt"

// Point struct
type Point struct {
	X, Y int
}

// Circle struct
type Circle struct {
	Point
	Radius int
}

// Wheel struct
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	fmt.Printf("%#v\n", w)

	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w2)
}
