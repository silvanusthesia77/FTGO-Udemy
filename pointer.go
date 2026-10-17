package main

import (
	"fmt"
	"strings"
)

type Adress struct {
	City, Province, Country string
}

func main() {
	adress1 := Adress{"Bandung", "Jawa barat", "Indonesia"}
	adress2 := adress1 //Copy Dari Adress 1
	adress2.City = "Jayapura"

	fmt.Println(adress1) // tidak Berubah
	fmt.Println(adress2) // hasil copy dari adress 1

	// Sekarang Kita Masuk Pada Pointer yang mana ini akan merubah semua data tidak mengcopy lagi.
	fmt.Println("\n", strings.Repeat("=", 45), "\n")
	adress3 := Adress{"Sorong", "Papua Barat Daya", "Indonesia"}

	adress4 := &adress3 // pointer &
	adress4.City = "Raja Ampat"

	fmt.Println(adress3) // Ikut berubah karena ini pointer &
	fmt.Println(adress4) // ikut berubah karena ini pointer &
}
