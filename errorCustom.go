package main

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n NotFoundError) Error() string {
	return n.Message
}
func ValidasiData(id string, data any) error {
	if id == "" {
		return &ValidationError{Message: "Data Error"}
	}
	if id != "Budi" {
		return &NotFoundError{Message: "Not Found"}
	} else {
		return nil
	}
}

// 54
