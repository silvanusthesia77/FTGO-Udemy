package main

import "fmt"

type students struct {
	name   string
	age    int
	adress string
}

func main() {
	var biodata students = students{"thoby", 20, "jln sagu"}
	var biodata1 *students = &biodata
	biodata1.name = "Biibbaa"

	fmt.Println(biodata)
	fmt.Println(biodata1)
}
