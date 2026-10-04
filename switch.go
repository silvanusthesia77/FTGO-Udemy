package main

import "fmt"

func main() {
	name := "Junior"

	switch name {
	case "Junior":
		fmt.Println("Hallo Junior")
	case "Alice":
		fmt.Println("Hallo Alice")
	default:
		fmt.Println("Boleh Kenalan ?")
	}

	switch length := len(name); length < 4 {
	case true:
		fmt.Println("Hallo Boleh Kenalan ?")
	case false:
		fmt.Println("Welcome", name)
	}
	name = "Parviddey"
	player := len(name)

	switch {
	case player > 10:
		fmt.Println("Nama terlalu panjang")
	case player < 5:
		fmt.Println("Nama terlalu Pendek")
	default:
		fmt.Println("Nama Sudah Benar", name)
	}
}
