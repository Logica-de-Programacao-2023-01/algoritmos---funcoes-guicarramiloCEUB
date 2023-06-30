package main

import (
	"errors"
	"fmt"
)

func filterStrings(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("slice is empty")
	}

	filteredSlice := make([]string, 0)

	for _, str := range slice {
		if len(str) > 5 {
			filteredSlice = append(filteredSlice, str)
		}
	}

	return filteredSlice, nil
}

func main() {
	slice := []string{"apple", "banana", "cherry", "durian", "elephant"}

	filteredSlice, err := filterStrings(slice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Original slice:", slice)
	fmt.Println("Filtered slice:", filteredSlice)
}
