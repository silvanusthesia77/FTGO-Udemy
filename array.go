package main

import (
	"fmt"
	"strings"
)

func main() {
	student := [3]string{}
	student[0] = "thb"
	student[1] = "Luiz"
	student[2] = "Arth"

	for _, v := range student {
		fmt.Println(v)
	}
	fmt.Println("\n", strings.Repeat("=", 38), "\n")
	sekolah := [3]string{
		"SD Cikombong",
		"SMP Cikombong",
		"SMA Abepura",
	}
	fmt.Println(sekolah[0])
	fmt.Println(sekolah[1])
	fmt.Println(sekolah[2])

	sports := [...]string{"Football", "tennis", "swimmin", "Running"}

	fmt.Println(sports)
	fmt.Println(len(sports))
}
