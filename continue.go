package main

import "fmt"

func main() {
	for j := 0; j < 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println("Hasil Continue :", j)
	}
}
