package main

import "fmt"

type Filter func(string) string

func sayHell(name string, filter Filter) {
	filterName := filter(name)
	fmt.Println("Hiii,", filterName)
}
func spanFilter(name string) string {
	if name == "Anjing" {
		return "......."
	} else {
		return (name)
	}
}
func main() {
	sayHell("Luizz Kandmy", spanFilter)
}

// anonymous Function
