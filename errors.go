package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("validation error")
	notFoundError   = errors.New("data not found")
)

func getById(id string) error {
	if id == "" {
		return validationError
	}
	if id != "eko" {
		return notFoundError
	}
	return nil

}

func main() {
	err := getById("")

	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, notFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
