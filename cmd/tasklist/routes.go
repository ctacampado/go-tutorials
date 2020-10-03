package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Response struct
type Response struct {
	Msg  string
	Data string
}

func createRsp(msg string, data []byte) ([]byte, error) {
	rsp := Response{Msg: msg, Data: string(data)}
	rbytes, err := json.Marshal(rsp)
	if err != nil {
		return nil, err
	}
	return rbytes, err
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rbytes, data []byte
	var err error
	var rsp Response

	data, err = getTasks()
	if err != nil {
		goto error
	}

	rsp = Response{Msg: "ok", Data: string(data)}
	rbytes, err = json.Marshal(rsp)
	if err != nil {
		goto error
	}
	w.Write(rbytes)
	return

error:
	fmt.Println("error:", err)
	w.Write([]byte(err.Error()))
}

func handleTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rbytes, data []byte
	var err error

	switch r.Method {
	case http.MethodGet:
		data, err = getTasks()
		if err != nil {
			goto error
		}
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			goto error
		}
		data = addTask(body)
	}

	rbytes, err = createRsp("ok", data)
	if err != nil {
		goto error
	}
	w.Write(rbytes)
	return

error:
	fmt.Println("error:", err)
	w.Write([]byte(err.Error()))
}
