package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 2
	d := 3
	e := a + b - c*d
	fmt.Println("Hasil :", e)

	i := 10
	i += 10
	fmt.Println("Nilai i :", i)
	g := 15
	g = g + 15
	fmt.Println("Nilai G :", g)
	j := 1
	j++
	j--
	fmt.Println("Hasil J :", j)

}
