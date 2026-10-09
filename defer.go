package main

import "fmt"

func login() {
	fmt.Println("Selesai Function")
}
func register() {
	defer login()
	fmt.Println("Welcome")
}
func main() {
	register()
}
