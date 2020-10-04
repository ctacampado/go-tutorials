package main

import (
	"go-tutorials/internal/kvstore"
	"net/http"
	"os"
	"time"
)

func getPort() string {
	arguments := os.Args
	if len(arguments) == 1 {
		return "3001"
	}
	return arguments[1]
}

func initServer(p string, m *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:         ":" + p,
		Handler:      m,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
}

// TaskList is the local in-mem k/v store for tasks
var TaskList = kvstore.New()

func main() {

	m := http.NewServeMux()
	m.HandleFunc("/", handleDefault)
	m.HandleFunc("/task", handleTask)
	m.HandleFunc("/timeout", handleTimeout)

	s := initServer(getPort(), m)
	s.ListenAndServe()
}
