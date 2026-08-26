package main

import "fmt"

func main() {
	sayHi("thoby", "junior")
	sayHi("Luiz", "Arthur")
}

func sayHi(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}
