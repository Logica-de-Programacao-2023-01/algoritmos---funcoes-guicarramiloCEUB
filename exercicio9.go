package main

import (
	"errors"
	"fmt"
	"strings"
)

func extractWordsFromString(str string) ([]string, error) {
	if len(str) == 0 {
		return nil, errors.New("string is empty")
	}

	words := strings.Fields(str)

	return words, nil
}

func main() {
	text := "Hello, world! Welcome to Go programming."

	wordList, err := extractWordsFromString(text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Original text:", text)
	fmt.Println("Words:", wordList)
}
