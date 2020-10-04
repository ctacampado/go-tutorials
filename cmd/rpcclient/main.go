package main

import (
	"fmt"
	"go-tutorials/internal/arith"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost"+":3003")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &arith.Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// Asynchronous call
	quotient := new(arith.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	fmt.Println(quotient)
	<-divCall.Done
	fmt.Println(quotient)
	//replyCall := <-divCall.Done // will be equal to divCall
	//fmt.Println(replyCall.ServiceMethod)
}
