package main

import "fmt"

func f1(out <-chan int64, in chan<- int64) {
	fmt.Println(<-out)
	// out <- 0
	in <- 20
	// fmt.Println(<-in)
}

func f2(in <-chan int64) {
	fmt.Println(<-in)
	// in <- 30
}

func main() {
	o := make(chan int64, 1)
	i := make(chan int64, 2)

	o <- 10
	i <- 2

	f1(o, i)
	f2(i)
}
