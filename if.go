package main

import "fmt"

func main() {

	name := "jokowi"

	if name == "luiz" {
		fmt.Println("Welcome Luiz Arthur")
	} else if name == "joko" {
		fmt.Println("Selamat Datang Joko")
	} else {
		fmt.Println("Halo !! boleh kenalan , ")
	}
	if length := len(name); length > 5 {
		fmt.Println("Lebih Besar dari : 5")
	} else {
		fmt.Println("Kurang Dari : 5")
	}
}

// switch expression
