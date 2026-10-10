package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("Masih Error", message)
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
}
func main() {
	runApp(true)
	fmt.Println("Bibba Welcome")
}
