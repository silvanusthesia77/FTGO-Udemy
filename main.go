package main

import "fmt"

func main() {
	fmt.Println("Hallo")
	greet("Luiz", "Swimming")
	number()
}
func greet(name string, hobby string) {
	fmt.Println(name, hobby)
}
func number() {
	// tipe data number
	fmt.Println("Number", 1)
	fmt.Println("Number Ls = ", 3.5)
	// tipe bool
	fmt.Println("Benar :", true)
	fmt.Println("Salah :", false)
	// tipe data string
}
