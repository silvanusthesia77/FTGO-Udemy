package main

import "fmt"

func sayHelloFiltter(name string, filter func(string) string) {
	filterName := filter(name)
	fmt.Println("Hiii", filterName)
}
func spamFlter(name string) string {
	if name == "Anjing" {
		return ".........."
	} else {
		return name
	}
}

func main() {
	hasil := spamFlter
	sayHelloFiltter("Biba", hasil)

	sayHelloFiltter("Anjing", spamFlter)
}
