package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func sortStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("slice is empty")
	}

	// Ordena o slice em ordem alfabética
	sort.Strings(slice)

	// Concatena as strings separadas por espaço
	result := strings.Join(slice, " ")

	return result, nil
}

func main() {
	slice := []string{"banana", "apple", "cherry", "durian"}

	sortedString, err := sortStrings(slice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Original slice:", slice)
	fmt.Println("Sorted string:", sortedString)
}
