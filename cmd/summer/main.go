package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(months[0])
	fmt.Println(months[11])
	fmt.Println(Q2)
	fmt.Println(summer)
	fmt.Println("Q2 cap: " + strconv.Itoa(cap(Q2)) + " len: " + strconv.Itoa(len(Q2)))
	fmt.Println("Q2 cap: " + strconv.Itoa(cap(summer)) + " len: " + strconv.Itoa(len(summer)))

	fmt.Println()

	nums := []int{2, 1, 6, 89, 21, 2}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)
}
