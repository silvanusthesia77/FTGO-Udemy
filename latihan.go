package main

import "fmt"

type students struct {
	name   string
	age    int
	adress string
}

func changess(stud *students) {
	stud.name = "Thobby"
	stud.age = 20
	stud.adress = "JLN.."
}

func main() {
	var studdt students = students{}
	changess(&studdt)
	fmt.Println(studdt)
}
