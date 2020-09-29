package main

import (
	"go-tutorials/internal/geometry"
)

func main() {
	pointReg := []geometry.Point{
		//{id: "1a2a", X: 1, Y: 1},
		{X: 1, Y: 1},
		{X: 4, Y: 6},
	}
	// pointReg[0].id = "1a2a"
	pointReg[0].SetID("1a2a")
	pointReg[1].SetID("1s2s")

	println(pointReg[0].GetID())
	println(pointReg[1].GetID())
}
