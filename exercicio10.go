package main

import (
	"errors"
	"fmt"
	"sort"
)

func sortSlice(numbers []int) ([]int, error) {
	if len(numbers) == 0 {
		return nil, errors.New("slice is empty")
	}

	sortedNumbers := make([]int, len(numbers))
	copy(sortedNumbers, numbers)

	sort.Ints(sortedNumbers)

	return sortedNumbers, nil
}

func main() {
	numbers := []int{5, 2, 9, 1, 7}

	sortedNumbers, err := sortSlice(numbers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Original numbers:", numbers)
	fmt.Println("Sorted numbers:", sortedNumbers)
}
