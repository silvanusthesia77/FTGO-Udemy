package main

import "fmt"

func main() {
	//var nilai int8 = 1
	//
	//for nilai <= 10 {
	//	fmt.Println("nilai :", nilai)
	//	nilai++
	//}
	//fmt.Println("Welldone")

	for counter := 1; counter < 10; counter++ {
		fmt.Println("nilai :", counter)
	}
	fmt.Println("Welldone")

	names := []string{"wanus", "thoby", "junior", "luiz"}
	// cara manual
	//for i := 0; i < len(names); i++ {
	//	fmt.Println(names[i])
	//}

	for index, value := range names {
		fmt.Println("Index", index, "=", "Value", value)
	}
}
