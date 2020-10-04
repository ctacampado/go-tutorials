package main

import (
	"go-tutorials/internal/arith"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(arith.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "localhost:3003")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}
