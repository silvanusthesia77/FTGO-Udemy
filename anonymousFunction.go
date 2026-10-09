package main

import "fmt"

type Filtter func(name string) bool

func filterfunck(name string, filtter Filtter) {
	if filtter(name) {
		fmt.Println("Have Blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	blocker := func(name string) bool {
		return name == "Anjing"
	}
	filterfunck("Thobby", blocker)
	filterfunck("Bibba", func(name string) bool {
		return name == "Anjing"
	})
}
