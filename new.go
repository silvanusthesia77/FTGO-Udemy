package main

import "fmt"

type Biodata struct {
	alamat string
	hobby  string
	usia   string
}

func main() {
	// var alamatt *Biodata = new(Biodata)  bisa menggunakan dibawah atau diatas juga sama saja
	var alamatt *Biodata = new(Biodata)
	var alamatt1 *Biodata = alamatt
	alamatt1.alamat = "Spain"
	fmt.Println(alamatt)
	fmt.Println(alamatt1)

	/* Cara Lain :

	almaat := new(Biodata)
	alamat1 := alamat
	alamat1.alamat = "Indonesia"

	*/

	// pages 48

}
