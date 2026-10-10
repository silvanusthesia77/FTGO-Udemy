package main

import "fmt"

type Customer struct {
	name, adress string
	age          int
}

func main() {
	var custmer Customer
	custmer.name = "Thoby"
	custmer.adress = "Km 17"
	custmer.age = 20

	fmt.Println(custmer)

	pelanggang := Customer{
		name:   "Wanus",
		adress: "Malibela",
		age:    22,
	}
	fmt.Println(pelanggang)

	budi := Customer{"Budi", "indonesia", 22}
	fmt.Println(budi)

	fmt.Println(Customer{"junior", "Japrax", 22})
}

//  pages 40
