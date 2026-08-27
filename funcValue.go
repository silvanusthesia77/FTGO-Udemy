package main

import "fmt"

func main() {
	group := terurName
	fmt.Println(group("Luiz"))
}
func terurName(names string) string {
	return "Hi, ... " + names
}

// function as parameter
