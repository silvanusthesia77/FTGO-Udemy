package main

import "fmt"

func fac(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * fac(value-1)
	}
}
func main() {
	fmt.Println(fac(10))
}

// 36
