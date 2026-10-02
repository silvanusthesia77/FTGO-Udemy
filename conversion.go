package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	fmt.Println(nilai32)
	var nilai64 int64 = int64(nilai32)
	fmt.Println(nilai64)
	var nilai16 int16 = int16(nilai32)
	fmt.Println(nilai16)

	var name = "thobiaz"
	var e = name[1]
	var string = string(e)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(string)
}

//Konversi tipe data 14
