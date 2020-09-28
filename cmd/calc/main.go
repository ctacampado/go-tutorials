package main

import "fmt"

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func main() {
	f := square
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)

	p := product
	fmt.Println(p(3, 4))
	fmt.Printf("%T\n", p)
	//f = product

	var g func(int) int
	g = square
	if g != nil {
		fmt.Println(g(6))
		fmt.Printf("%T\n", g)
	}
}
