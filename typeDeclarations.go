package main

import (
	"fmt"
	"strings"
)

func main() {
	type student string
	type noKtp string
	var nm string = "wanus"
	var ktp noKtp = "123456"
	var hasil noKtp = noKtp(nm)
	fmt.Println(hasil)
	fmt.Println(ktp)
	fmt.Println(noKtp("00065"))
	fmt.Println("\n", strings.Repeat("=", 30), "\n")
	var names student = "Luis Arthur"
	var value string = "00012345"
	var declaration student = student(names)
	fmt.Println(value)
	fmt.Println(declaration)

}

// Operasi Matematika
