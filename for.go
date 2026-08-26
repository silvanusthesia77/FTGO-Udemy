package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Hitung :", counter)
		counter++
	}
	fmt.Println("Selesai")

	for class := 2; class <= 8; class++ {
		fmt.Println("Thobby :", class)
	}
	fmt.Println("finish !!!")

	for Hasil := 0; Hasil < 6; Hasil++ {
		fmt.Println("Good Job", Hasil)
	}
	fmt.Println("Done")
	// cara manual
	names := []string{"wanus", "thoby", "luiz", "arthur"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Name :", names[i])
	}
	// cara otomatis

	students := []string{"Oliv", "Adella", "Defretes", "orpa", "luzi"}
	for index, name := range students {
		fmt.Println("Index :", index, "=", "Name :", name)
	}
}

// Break & Continue
