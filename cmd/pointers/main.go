package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(p)
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
	fmt.Println(incr(p))
	fmt.Println(incr(p))
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}
