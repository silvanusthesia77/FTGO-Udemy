package main

import "fmt"

func main() {
	var nilaiAkhir = 80
	var absensi = 90

	var lulusNilaiAkhir = nilaiAkhir > 70
	var lulusAbsensi = absensi > 70

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}
