package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func main() {
	PORT := ":3002"

	caCert, err := ioutil.ReadFile("./minica.pem")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	srv := &http.Server{
		Addr:      PORT,
		Handler:   &handler{},
		TLSConfig: cfg,
	}

	fmt.Println("Listening to port number", PORT)
	fmt.Println(srv.ListenAndServeTLS("./localhost/cert.pem", "./localhost/key.pem"))
}
