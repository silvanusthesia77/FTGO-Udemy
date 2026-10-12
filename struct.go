package main

import "fmt"

type Customer struct {
	name, adress string
	age          int
}

func (pelanggang Customer) sayHello(name string) {
	fmt.Println(name, "Hi, My name is :", pelanggang.name, "I am :", pelanggang.age, "My Adress is ", pelanggang.adress)
}

func main() {
	var customer Customer
	customer.name = "Bibba"
	customer.adress = "Pelangi"
	customer.age = 20
	fmt.Println(customer)

	var luiz = Customer{
		name:   "Luiz",
		adress: "Kilo 17",
		age:    22,
	}
	fmt.Println(luiz)
	fmt.Println(Customer{"thobby", "Pelangi", 22})
	luiz.sayHello("Andi")
}

//  41
