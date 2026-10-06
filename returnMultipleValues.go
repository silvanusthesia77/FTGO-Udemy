package main

import "fmt"

func deal() (string, string) {
	return "thoby", "Junior"
}
func main() {
	//firstname, lastname := deal()
	//fmt.Println(firstname, lastname)
	_, lastname := deal()
	fmt.Println(lastname)
}

//pages 30
