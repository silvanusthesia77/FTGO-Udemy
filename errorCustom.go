package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
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
		return &ValidationError{Message: "Validasi Error"}
	}
	if id != "EKo" {
		return &NotFoundError{Message: "Not Found Error"}
	}
	return nil
}
func main() {
	err := SaveData("", nil)
	if ValidationErr, ok := err.(*ValidationError); ok {
		fmt.Println(ValidationErr.Message)
	} else if NotFoundErr, ok := err.(*NotFoundError); ok {
		fmt.Println(NotFoundErr.Message)
	} else {
		fmt.Println(err.Error())
	}
}

// 54
