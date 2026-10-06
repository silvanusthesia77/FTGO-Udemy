package main

import "fmt"

func resault(name string) string {
	hasil := "Hii, " + name
	return hasil
}
func main() {
	var nama = resault("Harry")
	fmt.Println(nama)
	fmt.Println(resault("tHobby"))
}
