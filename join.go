package main

import "fmt"

func Join(arr []string, sep string) string {
	var r string
	for i := 0; i < len(arr); i++ {
		r += arr[i]
		if i != len(arr)-1 {
			r += sep
		}
	}
	return r
}

func main() {
	a := []string{"Aku", "Aku", "Delia"}
	fmt.Println(Join(a, "-"))
}
