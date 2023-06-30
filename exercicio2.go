package main

import "fmt"

func main() {
	s := "ceub"
	vogais(s)

}

func vogais(s string) {
	count := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "i" || string(s[i]) == "o" || string(s[i]) == "u" {
			count += 1
		}
	}
	fmt.Println(count)
}
