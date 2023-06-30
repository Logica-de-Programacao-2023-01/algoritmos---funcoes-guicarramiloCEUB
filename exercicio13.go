package main

import (
	"errors"
	"fmt"
)

func sumDigits(number int) (int, error) {
	if number < 0 {
		return 0, errors.New("number is negative")
	}

	sum := 0
	for number > 0 {
		digit := number % 10
		sum += digit
		number /= 10
	}

	return sum, nil
}

func main() {
	number := 12345

	sum, err := sumDigits(number)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Number: %d\nSum of digits: %d\n", number, sum)
}
