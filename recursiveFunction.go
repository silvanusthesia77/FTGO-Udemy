package main

import "fmt"

func factorial(value int) int {
	resault := 1
	for i := value; i > 0; i-- {
		resault *= i
	}
	return resault
}
func main() {
	fmt.Println(factorial(10))
}
