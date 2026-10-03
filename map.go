package main

import "fmt"

func main() {
	var biodata = map[string]string{
		"nama":  "Junior",
		"age":   "20",
		"hobby": "Football",
	}
	for _, v := range biodata {
		fmt.Println(v)
	}

	var person map[string]string = map[string]string{}
	person["name"] = "thobbiii"
	person["age"] = "22"

	fmt.Println(person["name"], person["age"])

	var mahasiswa = map[string]string{
		"name": "biiibbbaa",
		"age":  "22",
	}
	mahasiswa["name"] = "Dariuz"
	fmt.Println(mahasiswa["name"], mahasiswa["age"])

	var Universitas = make(map[string]string)
	Universitas["nama"] = "\"Universitas Siber Asia\", \"Universitas Bina Nusantara\""
	delete(Universitas, "nama")
	fmt.Println(Universitas["nama"])

	book := make(map[string]string)
	book["title"] = "Buku Go-lang"
	book["author"] = "wanus thesia"
	book["wrong"] = "ups"
	fmt.Println(book["title"], book["author"], book["wrong"])
	delete(book, "wrong")
	fmt.Println(book["title"], book["author"], book["wrong"])
}

//if expression 22
