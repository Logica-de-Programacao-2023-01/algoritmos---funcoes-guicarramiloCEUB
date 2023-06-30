package main

import "fmt"

func main() {
	s := []string{"oi", "ai", "ou"}
	fmt.Println(concat(&s))
}

func concat(s *[]string) string {
	final := ""
	for i := 0; i < len(*s); i++ {
		final += (*s)[i]
	}
	return final
}
