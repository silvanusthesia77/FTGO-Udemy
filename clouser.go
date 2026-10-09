package main

import "fmt"

func main() {
	counter := 0
	hasil := func() {
		fmt.Println("Bibba")
		counter++
	}

	hasil()
	hasil()
	hasil()
	fmt.Println(counter)
}
