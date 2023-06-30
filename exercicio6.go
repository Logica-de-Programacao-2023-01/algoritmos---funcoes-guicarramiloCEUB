package main

import "fmt"

func main() {
	s := []string{"oi", "ai", "ui"}
	concat2(s)
}

func concat2(s []string) {
	n := ""
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			n += s[i]
		} else {
			n += s[i] + ", "
		}
	}
	fmt.Println(n)
}
