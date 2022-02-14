package main

import "fmt"

func IsAnagram(s string, x string) {
	if len(s) != len(x) {
		fmt.Println("Not Anagram")
		return
	}

	mapS := make(map[string]int)

	for i := 0; i < len(s); i++ {
		mapS[string(s[i])]++
	}

	for i := 0; i < len(x); i++ {
		mapS[string(x[i])]--
	}

	for i := 0; i < len(s); i++ {
		if mapS[string(s[i])] != 0 {
			fmt.Println("Bukan Anagram")
			return
		}
	}

	fmt.Println("Anagram")
	return
}

func main() {
	IsAnagram("abc", "cba")
	IsAnagram("abc", "vbc")
	IsAnagram("danger", "garden")
}
