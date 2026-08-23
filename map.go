package main

import (
	"fmt"
)

func main() {
	person := make(map[string]string)
	person["Name"] = "wanus"
	person["hobby"] = "tidur"

	fmt.Println(person)

	student := map[string]string{
		"date": "223",
		"year": "1990",
		"not":  "Gila",
	}
	fmt.Println(student)

	delete(student, "not")
	fmt.Println(student)

}
