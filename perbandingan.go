package main

import (
	"fmt"
	"strings"
)

func main() {
	name1 := "thoby"
	name2 := "Thoby"
	fmt.Println(name1 == name2)
	fmt.Println(name1 != name2)

	fmt.Println("\n", strings.Repeat("=", 45), "\n")

}
