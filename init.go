package main

import (
	"FTGO-Udemy/database"
	_ "FTGO-Udemy/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDataBases())
}
