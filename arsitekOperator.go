package main

import (
	"fmt"
	"strings"
)

type Mahasiswa struct {
	name  string
	age   int
	hobby string
}

func main() {
	var biodata Mahasiswa = Mahasiswa{"Junior", 20, "Football"}
	var biodata1 = &biodata
	biodata1.hobby = "Fu**kk"
	// biodata1 = &Mahasiswa{"Bibbba", 21, "Player"}
	*biodata1 = Mahasiswa{"Luiz", 22, "Sulo"}
	fmt.Println(biodata)
	fmt.Println(biodata1)

	fmt.Println("\n", strings.Repeat("=", 35), "\n")

}
