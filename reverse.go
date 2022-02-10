package main

import "fmt"

func ReverseString(str string) string {
	var reversed string
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}

func ReverseArrayString(arr []string) []string {
	var reversed []string
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}

func main() {
	fmt.Println(ReverseString("Kodok"))
	fmt.Println(ReverseString("Hujan"))

	data := []string{"Kodok", "Hujan", "Kebon", "Angin"}
	fmt.Println(ReverseArrayString(data))
}
