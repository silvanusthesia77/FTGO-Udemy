package main

import "fmt"

func main() {
	var names = [3]string{"thoby", "Biba", "Parviddey"}
	fmt.Println(names)

	var name [3]string
	name[0] = "Luiz"
	name[1] = "arthur"
	name[2] = "Krish"
	name[2] = "thobiazz"
	fmt.Println(name)
	fmt.Println(len(names))
	for _, v := range names {
		fmt.Println(v)
	}

	var hasil = [3]int{}
	fmt.Println(hasil)
}

//tipe data slice 20
