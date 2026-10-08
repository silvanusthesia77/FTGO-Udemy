package main

import "fmt"

type FillterKataKotor func(name string) bool

func fiterKata(name string, filtter FillterKataKotor) {
	if name == "Anjing" {
		fmt.Println("Have Blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	fiterKata("Wanus", func(name string) bool {
		return name == "Anjing"
	})

	fill := func(name string) bool {
		return name == "Anjing"
	}
	fiterKata("wanus", fill)
}

//35
