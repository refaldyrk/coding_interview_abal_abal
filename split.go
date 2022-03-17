package main

import "fmt"

func Index(s, sep string) int {
	for i := 0; i < len(s); i++ {
		if s[i:i+len(sep)] == sep {
			return i
		}
	}
	return -1
}

func Split(s string, sep string) []string {
	var r []string
	for {
		i := Index(s, sep)
		if i == -1 {
			r = append(r, "Empty")
			break
		}
		r = append(r, s[:i])
		s = s[i+len(sep):]
	}
	return r
}

func main() {
	a := "Aku,Sayang,Delia"
	fmt.Println(Split(a, "-"))
}
