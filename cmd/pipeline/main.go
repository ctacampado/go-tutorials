package main

import (
	"fmt"
	"strconv"
)

func first(in chan string, out chan []int) {
	v, ok := <-in
	if ok && v == "count!" {
		fmt.Println(v)
		out <- []int{1}
	} else {
		out <- []int{}
	}
}

func second(in <-chan []int, out chan []int) {
	v, ok := <-in
	if ok && len(v) != 0 {
		fmt.Println(v)
		v = append(v, 2)
		out <- v
	} else {
		out <- []int{}
	}
}

func third(in <-chan []int, out chan<- []int) {
	v, ok := <-in
	if ok && len(v) != 0 {
		fmt.Println(v)
		v = append(v, 3)
		out <- v
	} else {
		out <- []int{}
	}
}

func fourth(in <-chan []int, out chan<- string) {
	v, ok := <-in
	if ok && len(v) == 3 {
		fmt.Println(v)
		data := ""
		for i, n := range v {
			data += strconv.Itoa(n)
			if i != len(v)-1 {
				data += "-"
			}
		}
		data += "-done!"
		out <- data
	} else {
		out <- ""
	}
}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan []int, 1)
	ch3 := make(chan []int, 1)
	ch4 := make(chan []int, 1)
	ch5 := make(chan string, 1)

	ch1 <- "count!"

	go first(ch1, ch2)
	go second(ch2, ch3)
	go third(ch3, ch4)
	go fourth(ch4, ch5)

	v, ok := <-ch5
	if ok && len(v) != 0 {
		fmt.Println(v)
	} else {
		fmt.Println("no data!")
	}
}
