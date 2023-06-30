package main

import (
	"errors"
	"fmt"
	"strings"
)

func replaceVowels(str string) (string, error) {
	if str == "" {
		return "", errors.New("string is empty")
	}

	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, vowel := range vowels {
		str = strings.ReplaceAll(str, vowel, "*")
	}

	return str, nil
}

func main() {
	str := "Hello, World!"

	result, err := replaceVowels(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Original string:", str)
	fmt.Println("Modified string:", result)
}
