package main

import "fmt"

func main() {
	ascii := 'a'
	unicode := 'ʒ'
	newline := '\n'
	korGo := '고'
	fmt.Printf("%d %[1]c %[1]q\n", ascii, unicode)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", korGo)
	fmt.Printf("%d %[1]q\n", newline)
}
