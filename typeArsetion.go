package main

import "fmt"

func roadMap() any {
	return true
}
func main() {
	hasil := roadMap()
	switch hasillroad := hasil.(type) {
	case string:
		fmt.Println("Hasil Dari String :", hasillroad)
	case int:
		fmt.Println("Hasil Dari Int :", hasillroad)
	default:
		fmt.Println("Unkwon", hasillroad)
	}
	// resault := pacc.(string)
	// fmt.Println(resault)
	// Dibawah adalah cara yang salah
	// resaultt := pacc.(int)
	// fmt.Println(resaultt)
}

// pages 45
