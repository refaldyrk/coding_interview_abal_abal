package main

import (
	"fmt"
	"sort"
)

func Set(arr []int) []int {
	var set []int
	for i := 0; i < len(arr); i++ {
		if !Contains(set, arr[i]) {
			set = append(set, arr[i])
		}
	}
	return set
}

func Contains(arr []int, n int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return true
		}
	}
	return false
}

func main() {
	data := []int{1, 1, 1, 3, 3, 4, 4, 5, 6, 4, 2, 2, 6}
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	fmt.Println(Set(data))
}
