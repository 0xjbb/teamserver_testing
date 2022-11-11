package main

type Listener struct {
	name string
	port string
}

func newListener() Listener {
	return Listener{}
}

func (*Listener) Spawn() {

}
