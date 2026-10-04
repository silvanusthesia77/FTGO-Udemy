package main

import "fmt"

func main() {
	name := "Thobiaz"

	if name == "Luiz" {
		fmt.Println("Selamat Datang", name)
	} else if name == "Thobiaz" {
		fmt.Println("Hi,", name)
	} else if name == "Junior" {
		fmt.Println("Hi,", name)
	} else {
		fmt.Println("Hi, Boleh Kenalan ?")
	}

	if length := len(name); length < 5 {
		fmt.Println("Name is too short")
	} else {
		fmt.Println(name)
	}
}
