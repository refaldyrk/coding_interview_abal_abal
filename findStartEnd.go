package main

import "fmt"

func FindStartEnd(arrInt []int, target int) []int {
	var r []int
	for i := 0; i < len(arrInt); i++ {
		if arrInt[i] == target {
			start := i
			for i+1 < len(arrInt) && arrInt[i+1] == target {
				i += 1
			}
			r = append(r, start, i)
		}
	}

	return r
}

func main() {
	data := []int{1, 4, 5, 5, 5, 5, 5, 6, 7} //2 - 6
	fmt.Println(FindStartEnd(data, 5))
}
