package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice": 31,
		"bob":   22,
	}

	fmt.Println(ages)
	for k, v := range ages {
		fmt.Printf("%s's age is %d\n", k, v)
	}

	ages["bob"]++
	fmt.Printf("bob's age after a year is %d\n", ages["bob"])

	fmt.Println(ages)
	fmt.Println(len(ages))
	delete(ages, "alice")
	fmt.Println(ages)
}
