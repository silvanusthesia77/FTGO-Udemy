package main

import (
	"fmt"
)

func main() {
	type noKTP string
	var idKTP noKTP = "111111"
	var KK string = "9999999"
	var idKK noKTP = noKTP(KK)

	fmt.Println("idKTP :", idKTP)
	fmt.Println("KK :", KK)
	fmt.Println("idKK :", idKK)

}

// Operasi Matematika
