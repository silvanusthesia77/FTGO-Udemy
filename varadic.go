package main

import "fmt"

func main() {
	number := sumAll(10, 20, 30, 40, 50)
	fmt.Println(number)

	// kalau semisal kita sudah punya variable
	hasill := []int{10, 10, 10, 10}
	fmt.Println(sumAll(hasill...))
}
func sumAll(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}
