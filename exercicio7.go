package main

import (
	"errors"
	"fmt"
)

func applyFunctionToSlice(numbers []int, fn func(int) int) ([]int, error) {
	if len(numbers) == 0 {
		return nil, errors.New("slice is empty")
	}

	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = fn(num)
	}

	return result, nil
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Exemplo de função que eleva um número ao quadrado
	square := func(num int) int {
		return num * num
	}

	// Aplica a função `square` em cada elemento do slice
	squaredNumbers, err := applyFunctionToSlice(numbers, square)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Original numbers:", numbers)
	fmt.Println("Squared numbers:", squaredNumbers)
}
