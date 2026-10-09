package main

import "fmt"

type FilterBlock func(name string) bool

func kataKotor(name string, fillter FilterBlock) {
	if fillter(name) {
		fmt.Println("Have Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
func main() {
	fill := func(name string) bool {
		return name == "Anjing"
	}
	kataKotor("thoby", fill)

	kataKotor("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
