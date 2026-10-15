package main

import "fmt"

func pack() any {
	return true
}
func main() {
	pacc := pack()
	// resault := pacc.(string)
	// fmt.Println(resault)
	// Dibawah adalah cara yang salah
	// resaultt := pacc.(int)
	// fmt.Println(resaultt)

	switch resault := pacc.(type) {
	case string:
		fmt.Println("Hasil String ", resault)
	case int:
		fmt.Println("Int Hasilnya :", resault)
	default:
		fmt.Println("Unknown", resault)
	}
}

// pages 45
