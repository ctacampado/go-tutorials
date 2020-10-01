package main

import (
	"fmt"
	"sync"
)

func printer(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s\n", <-c)
}

func main() {
	wg := sync.WaitGroup{}
	msg := "This message was sent via a channel!"
	//msg2 := "This message was also sent via a channel!"

	c := make(chan string, 1)
	c <- msg
	// fmt.Printf("%s\n", <-c)
	// c <- msg2
	//wg.Add(1)
	go printer(c, &wg)
	//wg.Wait()
	close(c)
	v, ok := <-c
	if ok {
		fmt.Println("Channel is still open and contains: [", v, "]")
	} else {
		fmt.Println("Channel is closed")
	}
}
