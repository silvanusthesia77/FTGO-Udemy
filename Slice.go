package main

import (
	"fmt"
	"strings"
)

func main() {
	names := [...]string{"biba", "luiz", "arthur", "Reza", "golix"}
	fmt.Println(names[1:])
	slice := names[:3]
	fmt.Println(slice)
	slice1 := names[:]
	fmt.Println(slice1)

	hobbies := []string{"Sport", "reading", "running", "selling"}
	hobbies[1] = "Membaca"
	hobbies = append(hobbies, "My Love Siiihhh")
	fmt.Println(hobbies)

	fmt.Println("\n", strings.Repeat("=", 35), "\n")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "thby"
	newSlice[1] = "arthur"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice1 := append(newSlice, "Galau")
	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))

	fmt.Println("\n", strings.Repeat("=", 36), "\n")

	fromSlice := hobbies[:]
	toClice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toClice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toClice)
}

// Tipe Data Map
