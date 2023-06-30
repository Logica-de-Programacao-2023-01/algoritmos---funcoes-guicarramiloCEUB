package main

import (
	"errors"
	"fmt"
)

func applyAndSum(slice []int, fn func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("slice is empty")
	}

	sum := 0
	for _, num := range slice {
		result := fn(num)
		sum += result
	}

	return sum, nil
}

func double(n int) int {
	return n * 2
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	result, err := applyAndSum(slice, double)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Slice:", slice)
	fmt.Println("Result:", result)
}
