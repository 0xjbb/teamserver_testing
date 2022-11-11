package main

import "fmt"

wg := &sync.WaitGroup



func main() {

	test8080 := newListener("8080")
	test8080.Spawn()

	fmt.Println("blocking")

	test8888 := newListener("8888")
	test8888.Spawn()
}
