package main

import (
	"errors"
	"fmt"
)

func containsNumber(number int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("slice is empty")
	}

	for _, num := range slice {
		if num == number {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	number := 7
	slice := []int{1, 3, 5, 7, 9}

	contains, err := containsNumber(number, slice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Slice:", slice)
	fmt.Printf("Number %d is present: %v\n", number, contains)
}
