package main

import "fmt"

func main() {
	threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}
	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}
	}

	fmt.Println()

	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
	}

	fmt.Println()
}
