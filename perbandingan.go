package main

import "fmt"

func main() {
	var name1 = "Thoby"
	var name2 = "Junior"
	var resault = name1 == name2
	fmt.Println(resault)
	fmt.Println(name1 != name2)

	a := 10
	b := 2
	c := a > b
	fmt.Println(c)
	fmt.Println(a < b)
}
