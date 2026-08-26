package main

import "fmt"

func main() {
	// cara pertama
	// gett := getHello("tHoby")
	// fmt.Println(gett)
	// cara ke-2
	fmt.Println(getHello("Biba"))
}
func getHello(name string) string {
	get := "Hii, Saya " + name
	return get
}

// Returning Multiple Values
