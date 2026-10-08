package main

import "fmt"

func sayHi(name string) string {
	return "Hi " + name
}
func main() {
	gethi := sayHi
	fmt.Println(gethi("Luizzz"))
}
