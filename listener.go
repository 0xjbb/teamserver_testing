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
	http.HandleFunc("/", lst.HandleSlash)
	http.ListenAndServe(":"+lst.Port, nil)
}

func (lst *Listener) HandleSlash(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This a teamserver arch test! On port "+lst.Port+"\n")
}
