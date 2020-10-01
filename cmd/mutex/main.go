package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m sync.Mutex

// STOP determines if gouroutine stageA will have stop or continue
var STOP = false

// DATA is an integer map who's key is also the value
var DATA = make(map[int]int)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func stageA(min, max int, out chan<- int, wg *sync.WaitGroup) {
	for {
		m.Lock()
		if STOP {
			m.Unlock()
			return
		}
		m.Unlock()
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
			m.Lock()
			STOP = true
			m.Unlock()
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

	close(ch1)
	v, ok := <-ch1
	if ok {
		fmt.Println("open!", v)
	} else {
		fmt.Println("close")
	}

	fmt.Println(DATA)
}
