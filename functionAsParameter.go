package main

import "fmt"

type Filter func(string) string

func filterkatakotor(name string, filter Filter) {
	filtername := filter(name)
	fmt.Println("Hiii, ", filtername)
}
func filter(name string) string {
	if name == "Anjing" {
		return "......"
	} else {
		return name
	}
}
func main() {
	filterr := filter
	filterkatakotor("Anjing", filterr)
	filterkatakotor("Thobiazz ", filter)
}

//34 anonymous function
