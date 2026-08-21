package main

import "fmt"

func main() {
	var nilai32 int32 = 32767
	var nilai16 int16 = int16(nilai32)
	var nilai64 int64 = int64(nilai16)

	fmt.Println("Nilai32 : ", nilai32)
	fmt.Println("Nilai16 : ", nilai16)
	fmt.Println("Nilai64 : ", nilai64)

	nm := "darling"
	var e uint8 = nm[1]
	var estring = string(e)
	fmt.Println(nm)
	fmt.Println(e)
	fmt.Println(estring)
}

// Type Declarations
