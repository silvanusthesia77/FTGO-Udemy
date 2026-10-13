package main

import "fmt"

func sayHellos(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Nama": name,
		}
	}
}
func main() {
	data := sayHellos("")
	if data == nil {
		fmt.Println("Data Masih Kosong")
	} else {
		fmt.Println(data["Nama"])
	}
}

//  page 43 nil
