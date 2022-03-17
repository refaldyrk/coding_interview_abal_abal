package main

import "fmt"

func Count(s string, c string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c[0] {
			fmt.Println(string(s[i]))
			count++
		}
	}
	return count
}

func main() {
	word := "Aku Sayang Delia, Sayang Banget, Kamu Sayang Tidak?"

	fmt.Println(Count(word, "S"))
}
