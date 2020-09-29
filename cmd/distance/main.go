package main

import (
	"fmt"
	"go-tutorials/internal/geometry"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}

	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))

	fmt.Println()

	tri := geometry.Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}

	fmt.Println(geometry.PathDistance(tri))
	fmt.Println(tri.Distance())

	fmt.Println()

	p.UpdateX(3)
	fmt.Println(p.Distance(q))
	fmt.Println(p)
	p.ScalyBy(2)
	fmt.Println(p)
}
