package main

import (
	"io"
	"net/http"
)

func main() {
	/*
		test8080 := newListener("8080")
		test8080.Spawn()

		fmt.Println("blocking")

		test8888 := newListener("8888")
		test8888.Spawn()
	*/

	server := http.NewServeMux()
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "This the main dashbaord area \n")

	})

	server.HandleFunc("/listener_8888", func(w http.ResponseWriter, r *http.Request) {
		test8888 := newListener("8888")
		test8888.Spawn()
	})

	server.HandleFunc("/listener_8080", func(w http.ResponseWriter, r *http.Request) {
		test8888 := newListener("8080")
		test8888.Spawn()
	})

	http.ListenAndServe(":42069", server)
}
