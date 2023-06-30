package main

import (
	"errors"
	"fmt"
	"math"
)

func isPrimeNumber(n int) (bool, error) {
	if n < 0 {
		return false, errors.New("number is negative")
	}

	if n < 2 {
		return false, nil
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	number := 17

	isPrime, err := isPrimeNumber(number)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Number: %d\nIs Prime: %v\n", number, isPrime)
}
