package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	var addrss *Address = &Address{}
	changeCountryToIndonesia(addrss)
	fmt.Println(addrss)
}
