package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	h := squares()
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
	fmt.Println(h())
}
