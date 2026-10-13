package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasname HasName) {
	fmt.Println(hasname.GetName())
}

type Person struct {
	name   string
	age    int
	adress string
}

func (person Person) GetName() string {
	return person.name
}
func main() {
	biodata := Person{name: "Wanus", age: 20, adress: "kilo 18"}
	fmt.Println(biodata)
}
