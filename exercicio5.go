package main

import "fmt"

func main() {
	n := 4
	s := []int{1, 4, 3, 4, 4, 5}
	iguais(&s, n)
}

func iguais(s *[]int, n int) {
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == n {
			fmt.Println(i + 1)
			break
		}
	}
}
