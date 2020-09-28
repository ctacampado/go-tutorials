package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1234612309834
	nS := strconv.Itoa(n)
	fmt.Println(comma(nS))
}

// takes a string representation of an integer and insert
// commas every three places:
// 12345 -> 12,345
func comma(s string) string {
	numStr := s
	n := len(numStr)
	for i := n; i >= 0; i-- {
		if i-2 > 0 && i-3 != 0 {
			numStr = numStr[:i-3] + "," + numStr[i-3:]
			i -= 2
		}
	}
	return numStr
}

// same as comma but for floating-points
func commaFloat(s string) string {
	var fs string
	// implementation here
	return fs
}
