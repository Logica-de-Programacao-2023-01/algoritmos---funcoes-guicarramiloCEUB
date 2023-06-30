package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func filterUpperCaseStrings(stringsSlice []string) (string, error) {
	if len(stringsSlice) == 0 {
		return "", errors.New("slice is empty")
	}

	var filteredStrings []string
	for _, str := range stringsSlice {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			filteredStrings = append(filteredStrings, str)
		}
	}

	result := strings.Join(filteredStrings, " ")

	return result, nil
}

func main() {
	stringsSlice := []string{"Hello", "world", "Welcome", "to", "Go"}

	filteredString, err := filterUpperCaseStrings(stringsSlice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Original strings:", stringsSlice)
	fmt.Println("Filtered string:", filteredString)
}
