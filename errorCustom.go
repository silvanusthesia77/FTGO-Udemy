package main

import "fmt"

type ValidasiError struct {
	Message string
}

func (v *ValidasiError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidasiError{Message: "Validasi Error"}
	} else if id != "Andy" {
		return &NotFoundError{Message: "Not Found Error"}
	} else {
		return nil
	}

}
func main() {
	err := SaveData("And", nil)
	if err != nil {
		// if ValidasiErr, ok := err.(*ValidasiError); ok {
		// 	fmt.Println(ValidasiErr.Error())
		// } else if NotFoundErr, ok := err.(*NotFoundError); ok {
		// 	fmt.Println(NotFoundErr.Error())
		// } else {
		// 	fmt.Println("Hasil Akhir :", err.Error())
		// }
		switch finalError := err.(type) {
		case *ValidasiError:
			fmt.Println(finalError.Error())
		case *NotFoundError:
			fmt.Println(finalError.Error())
		default:
			fmt.Println(finalError.Error())
		}
	} else {
		fmt.Println("Suksess")
	}
}

// 55 Materi selanjutnya
