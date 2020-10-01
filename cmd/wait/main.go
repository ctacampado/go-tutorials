package main

import (
	"fmt"
	"time"
)

func main() {
	n := 999
	//wg := sync.WaitGroup{}
	fmt.Printf("Creating %d goroutines...\n", n)

	for i := 0; i < n; i++ {
		//wg.Add(1)
		go func(i int) {
			//defer wg.Done()
			fmt.Printf("%d ", i)
		}(i)
		// wg.Done() // negative waitgroup
	}
	// wg.Add(1) // deadlock

	time.Sleep(time.Millisecond)
	//wg.Wait()
	fmt.Printf("\nExiting...\n")
}
