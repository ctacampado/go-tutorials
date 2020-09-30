package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var closing uint32 = 0

// DATA is an integer map who's key is also the value
var DATA = make(map[int]int)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func stageA(min, max int, out chan<- int, wg *sync.WaitGroup) {
	for {
		if atomic.LoadUint32(&closing) != 0 {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func stageB(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var i int
	for {
		i = <-in
		if DATA[i] != 0 {
			fmt.Println(DATA[i], i)
			atomic.StoreUint32(&closing, 1)
			return
		}
		DATA[i] = i
	}
}

func main() {
	ch1 := make(chan int)
	wg := sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())
	wg.Add(1)
	go stageA(1, 50, ch1, &wg)
	go stageB(ch1, &wg)
	wg.Wait()
	fmt.Println(DATA)
}
