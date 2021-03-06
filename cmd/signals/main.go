package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type signal struct{}

// DATA is an integer map who's key is also the value
var DATA = make(map[int]int)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func stageA(min, max int, out chan<- int, stopSig <-chan signal, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopSig:
			return
		case out <- random(min, max):
		}
	}
}

func stageB(in <-chan int, stopSig chan<- signal, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		i := <-in
		if DATA[i] != 0 {
			fmt.Println(DATA[i], i)
			stopSig <- signal{}
			return
		}
		DATA[i] = i
	}
}

func main() {

	sigc := make(chan signal)
	ch1 := make(chan int)

	rand.Seed(time.Now().UnixNano())

	wg := sync.WaitGroup{}
	wg.Add(2)
	go stageA(1, 50, ch1, sigc, &wg)
	go stageB(ch1, sigc, &wg)

	wg.Wait()

	close(ch1)
	v, ok := <-ch1
	if ok {
		fmt.Println("open!", v)
	} else {
		fmt.Println("close")
	}

	fmt.Println(DATA)
}
