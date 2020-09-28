package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	s1 := "cc1"
	s2 := "1cc"

	if isAnagram(s1, s2) {
		fmt.Printf("%s is an anagram of %s!\n", s1, s2)
	} else {
		fmt.Printf("%s is NOT an anagram of %s!\n", s1, s2)
	}
}

func isAnagram(s1, s2 string) bool {
	var hits []byte

	if len(s1) != len(s2) {
		return false
	}

	for _, v1 := range s1 {
		for i, v2 := range s2 {
			idx := strconv.Itoa(i)
			if v1 == v2 && !bytes.Contains(hits, []byte(idx)) {
				hits = bytes.Join([][]byte{hits, []byte(idx)}, []byte(""))
				if len(hits) == len(s1) {
					return true
				}
				break
			}
			if i == len(s2)-1 {
				return false
			}
		}
	}
	return false
}
