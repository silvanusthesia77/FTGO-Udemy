package main

import "fmt"

type noKTP string

func main() {
	var a noKTP = "123456"
	fmt.Println(a)
	var b string = ("000999666")
	var c noKTP = noKTP(b)
	fmt.Println(c)

}
