package main

import "fmt"

func sayHelloFilter(name string, filter func(string) string) {
	spanName := filter(name)
	fmt.Println("Hi,.........", spanName)
}
func spamFillter(name string) string {
	if name == "Anjing" {
		return (".........")
	} else {
		return name

	}
}
func main() {
	filtter := spamFillter
	sayHelloFilter("Luizzz", filtter)
}
