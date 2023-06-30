package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(pares(s))
}

func pares(s []int) []int {
	n := []int{}
	if len(s) == 0 {
		return nil
	}
	for i := 0; i < len(s); i++ {
		if s[i]%2 == 0 {
			n = append(n, s[i])
		}
	}
	return n
}
