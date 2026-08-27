package main

import "fmt"

func main() {
	// ini kalau kita butuh semuanya
	// firstName, lastName := getFullname()
	// fmt.Println(firstName, lastName)

	// ini kalau kita hanya butuh FirstName
	firstName, _ := getFullname()
	fmt.Println(firstName)
}

func getFullname() (string, string) {
	return "tHoby", "Junior"
}
