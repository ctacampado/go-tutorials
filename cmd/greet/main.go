package main

import "fmt"

func main() {
	s := "Hello, world"
	hello := s[:5]
	world := s[7:]

	fmt.Printf("%s %d\n", s, len(s))
	fmt.Printf("%s %d\n", hello, len(hello))
	fmt.Printf("%s %d\n", world, len(world))

	s = "Hello, 세계"
	hello = s[:5]
	world = s[7:]

	fmt.Printf("%s %d\n", s, len(s))
	fmt.Printf("%s %d\n", hello, len(hello))
	fmt.Printf("%s | % x | len: %d\n", world, world, len(world))
}
