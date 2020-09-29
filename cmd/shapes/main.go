package main

import (
	"go-tutorials/internal/shape"
)

func main() {
	square := shape.Square{X: 4}
	circle := shape.Circle{R: 2}

	shape.Calculate(circle)
	shape.Calculate(square)
}
