package main

import "fmt"

// Weekday type
type Weekday int

const (
	// Sunday = 0
	Sunday Weekday = iota
	// Monday = 1
	Monday
	// Tuesday = 2
	Tuesday
	// Wednesday = 3
	Wednesday
	// Thursday = 4
	Thursday
	// Friday = 5
	Friday
	// Saturday = 6
	Saturday
)

func main() {
	days := []Weekday{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}
	for _, v := range days {
		fmt.Println(v)
	}
}
