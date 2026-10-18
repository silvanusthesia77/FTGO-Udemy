package main

import "fmt"

type Nam struct {
	name string
}

func (nam *Nam) Canges() {
	nam.name = "Mr. " + nam.name
}
func main() {
	var nama Nam = Nam{"Thobby"}
	nama.Canges()
	fmt.Println(nama.name)
}

// pages 50
