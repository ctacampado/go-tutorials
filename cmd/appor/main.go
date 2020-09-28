package main

import "fmt"

func main() {
	var apples int32 = 1
	var oranges int16 = 2
	// var compote int = apples + oranges
	var compote = int(apples) + int(oranges)
	fmt.Println(compote)
}
