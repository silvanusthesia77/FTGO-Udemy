package main

import "fmt"

func main() {
	names := [...]string{"thoby", "junior", "parveddey", "biba", "reza", "luiz"}
	slice := names[0:2]
	slice2 := names[1:3]
	slice = append(slice, names[1])
	fmt.Println(slice)
	fmt.Println(len(slice2))

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	dayClice1 := days[5:]
	dayClice1[0] = "Libur Buruh"
	dayClice1[1] = "Libur Panjang"
	fmt.Println(days)
	dayClice2 := append(dayClice1, "Libur Nasional")
	fmt.Println((dayClice2))
	fmt.Println(cap(dayClice1))
	dayClice3 := copy(dayClice1, dayClice2)
	fmt.Println(dayClice3)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "wanus"
	newSlice[1] = "junior"
	//newSlice[2] = "parveddey" error karena sudah dibatasi hanya 2
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	var newSlice1 = append(newSlice, "tHobiazz")
	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))

	newSlice1[0] = "Bibbaa"
	fmt.Println(newSlice1)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
	fmt.Println(fromSlice)

	slices := []int{1, 2, 3, 4, 5}
	arrays := [...]int{1, 2, 3, 4, 5}
	fmt.Println(slices)
	fmt.Println(arrays)

	var sekolah = make([]string, 3, 5)
	sekolah[0] = "Libur Buruh"
	sekolah[1] = "Libur Panjang"
	fmt.Println(sekolah)
	fmt.Println(len(sekolah))
	fmt.Println(cap(sekolah))
	tambah := append(sekolah, "Hari ini jumat")
	fmt.Println(tambah)
	fmt.Println(len(tambah))
	fmt.Println(cap(tambah))
}

//tipe data map 21
