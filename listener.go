package main

import (
	"io"
	"net/http"
)

type Listener struct {
	Name string
	Port string
}

func newListener(port string) Listener {
	return Listener{Port: port}
}

func (lst *Listener) Spawn() {
	testMux := http.NewServeMux()
	testMux.HandleFunc("/", lst.HandleSlash)
	go http.ListenAndServe(":"+lst.Port, testMux)
}

func (lst *Listener) HandleSlash(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This a teamserver arch test! On port "+lst.Port+"\n")
}
