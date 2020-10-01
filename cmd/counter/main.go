package main

import (
	"fmt"
	"time"
)

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

func main() {

	go function()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()

	/*
		for i := 0; i < 100; i++ {
			go func(x int) {
				fmt.Printf("%d ", x)
			}(i)
		}
	*/

	time.Sleep(time.Second)
	fmt.Println()
}
