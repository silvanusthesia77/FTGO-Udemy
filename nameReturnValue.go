package main

import "fmt"

func getFullName() (name, username, adress string) {
	name = "John Doe"
	username = "jdoe"
	adress = "Jakarta"
	return name, username, adress
}
func main() {
	name, username, adress := getFullName()
	fmt.Println(name, username, adress)
}
