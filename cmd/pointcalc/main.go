package main

import (
	"fmt"
	"go-tutorials/internal/geometry"
)

func main() {
	tri := geometry.Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}
	fmt.Printf("%+v\n", tri)
	offset := geometry.Point{X: 2, Y: 2}
	tri.TranslateBy(offset, true)
	fmt.Printf("%+v\n", tri)
}
