package main

import "fmt"

type BlochName func(name string) bool

func blockName(name string, block BlochName) {
	if name == "anjing" {
		fmt.Println("Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
func main() {
	hasil := func(name string) bool {
		return name == "anjing"
	}
	blockName("anjing", hasil)
}

// recursive function
