package main

import "fmt"

func main() {
	var name string
	var age int
	name = "Bob"
	age = 25
	fmt.Println(name, age)
	name = "Biba"
	age = 20
	fmt.Println(name, age)

	var hobby = "Sport"
	fmt.Println(hobby)

	address := "Jln Dukuh"
	fmt.Println(address)
	address = "Jln Kalibata"
	fmt.Println(address)

	var (
		firstName = "Junior"
		lastName  = "BIba"
		hobbies   = "Play Football"
		_         = hobbies
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
