package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Response struct
type Response struct {
	Msg    string
	Data   string
	Status int
}

func createRsp(msg string, data []byte, err error) Response {
	var rsp Response
	rsp.Msg = msg
	if err != nil {
		rsp.Data = err.Error()
		rsp.Status = http.StatusInternalServerError
	} else {
		rsp.Data = string(data)
		rsp.Status = http.StatusOK
	}
	return rsp
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	var rbytes, data []byte
	var err error
	var rsp Response

	msg := "ok"

	data, err = getTasks()
	if err != nil {
		msg = "error"
	}

	rsp = createRsp(msg, data, err)
	rbytes, err = json.Marshal(rsp)
	if err != nil {
		rsp = createRsp("error", nil, err)
	}

	w.WriteHeader(rsp.Status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rbytes)
}

func handleTask(w http.ResponseWriter, r *http.Request) {
	var rbytes, data []byte
	var err error
	var rsp Response

	msg := "ok"

	switch r.Method {
	case http.MethodGet:
		data, err = getTasks()
		if err != nil {
			msg = "error"
		}
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			msg = "error"
		}
		data = addTask(body)
	}

	rsp = createRsp(msg, data, err)
	rbytes, err = json.Marshal(rsp)
	if err != nil {
		rsp = createRsp("error", nil, err)
	}

	w.WriteHeader(rsp.Status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rbytes)
}

func handleTimeout(w http.ResponseWriter, r *http.Request) {
	var rbytes []byte
	var err error
	var rsp Response

	//time.Sleep(2 * time.Second)
	rsp = createRsp("ok", []byte("no delay"), nil)
	rbytes, err = json.Marshal(rsp)
	if err != nil {
		rsp = createRsp("error", nil, err)
	}

	w.WriteHeader(rsp.Status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rbytes)
}
