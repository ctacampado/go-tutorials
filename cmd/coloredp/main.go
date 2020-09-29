package main

import (
	"fmt"
	"go-tutorials/internal/geometry"
	"image/color"
)

// ColoredPoint represents a colored point
type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}

// Distance method wrapper
// func (p *ColoredPoint) Distance(q geometry.Point) float64 {
// 	return p.Point.Distance(q)
// }

// ScaleBy method wrapper
// func (p *ColoredPoint) ScaleBy(factor float64) {
//	p.Point.ScalyBy(factor)
// }

func main() {
	var cp ColoredPoint
	cp.X = 1
	cp.Point.Y = 2
	fmt.Println(cp.Point.X, cp.Y)
	fmt.Printf("%+v\n", cp)

	fmt.Println()

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint{
		geometry.Point{X: 1, Y: 1},
		red,
	}
	fmt.Printf("%+v\n", p)
	q := ColoredPoint{
		geometry.Point{X: 5, Y: 4},
		blue,
	}
	fmt.Printf("%+v\n", q)
	fmt.Println("distance between p & q: ", p.Distance(q.Point))
	q.ScalyBy(2)
	fmt.Printf("%+v\n", q)
	// p.Distance(q)
}
