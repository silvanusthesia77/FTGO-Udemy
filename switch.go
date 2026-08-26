package main

import "fmt"

func main() {
	name := "Thobiii"

	switch {
	case name == "Luiz":
		fmt.Println("Selamat , Datang Luizz !!! ")
	case name == "Thoby":
		fmt.Println("Selamat Datang , Biba !!")
	default:
		fmt.Println("Hi, Boleh Kenalan ?")
	}

	length := len(name)
	switch {
	case length > 8:
		fmt.Println("Selamat, Datang KembLI !!!")
	case length <= 6 && length > 4:
		fmt.Println("Lumayan")
	default:
		fmt.Println("Back Again !!")
	}

	switch lengt := len(name); lengt > 5 {
	case true:
		fmt.Println("Nama terlalu panjan")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
