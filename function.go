package main

import "fmt"

func sayHello() {
	fmt.Println("Hi, Selamat datang yeah ...!!!")
}
func student(name string, age int) {
	fmt.Println("Name :", name, "=", "Age :", age)
}
func main() {
	sayHello()
	student("Luiz Arthur", 22)
}
