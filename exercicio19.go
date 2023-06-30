package main

import (
	"errors"
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primeNumbers(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("negative number")
	}

	primes := make([]int, 0)

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes, nil
}

func main() {
	number := 20

	primeSlice, err := primeNumbers(number)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Number:", number)
	fmt.Println("Prime numbers:", primeSlice)
}