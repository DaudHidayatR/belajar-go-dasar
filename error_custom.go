package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message

}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message

}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error: id is required"}
	}
	if id != "daud" {
		return &notFoundError{"data not found"}
	}
	return nil
}

func main() {
	err := SaveData("daud", nil)
	if err != nil {
		switch v := err.(type) {
		case *validationError:
			fmt.Println("validation error:", v.Message)
		case *notFoundError:
			fmt.Println("not found error:", v.Message)
		default:
			fmt.Println("unknown error")
		}
	} else {
		fmt.Println("data berhasil disimpan")
	}
}
