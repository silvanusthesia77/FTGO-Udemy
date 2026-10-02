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

	var data1 int32 = 32768
	fmt.Println(data1)
	var data2 int64 = int64(data1)
	fmt.Println(data2)
	var data3 int8 = int8(data1)
	fmt.Println(data3)
}

//Konversi tipe data 14
