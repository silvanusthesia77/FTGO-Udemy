package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func (cusyomer Customer) sayHello(name string) {
	fmt.Println(name, "Hi, I am ", cusyomer.name, "I live in ", cusyomer.address, "I am ", cusyomer.age)
}

func main() {
	var andy = Customer{
		name:    "Andy",
		address: "Jln. Ambrobem",
		age:     20,
	}
	fmt.Println(andy)

	var Luiz Customer
	Luiz.name = "Endy"
	Luiz.address = "Kilo 17"
	Luiz.age = 22

	fmt.Println(Luiz)
	Luiz.sayHello("Andy")
}
