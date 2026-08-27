package main

import "fmt"

func main() {
	a, b, c := getName()
	fmt.Println(a, b, c)
}

func getName() (firstName, middleName, lastName string) {
	firstName, middleName, lastName = "thobias", "parviddey", "junior"
	return firstName, middleName, lastName
}

// variadic function
