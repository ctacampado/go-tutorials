package main

import (
	"fmt"
	"time"
)

// Rocket struct
type Rocket struct {
	Message string
}

// Launch prints a message
func (r *Rocket) Launch() {
	fmt.Println(r.Message)
}

func main() {
	r := Rocket{"Bye, World!"}
	time.AfterFunc(5*time.Second, r.Launch)
	time.Sleep(6 * time.Second)
}
