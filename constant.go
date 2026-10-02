package main

import (
	"fmt"
)

func main() {
	const name string = "Luiz"
	fmt.Println(name)
	const hobby = "Sport"
	fmt.Println(hobby)
	//name = "Jonu" tidak dapat di assign

	const (
		firstName = "James"
		lastName  = "Bond"
	)
	fmt.Println(firstName, lastName)

}
