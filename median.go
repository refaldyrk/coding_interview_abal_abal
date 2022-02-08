package main

import (
	"fmt"
	"sort"
)

func Median(data []int) float64 {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	middleData := len(data) / 2

	if len(data)%2 == 0 {
		calcMedian := (data[middleData-1] + data[middleData]) / 2
		return float64(calcMedian)
	}

	return float64(data[middleData])
}

func main() {
	data := []int{3, 4, 5, 5, 6, 6, 7, 7, 7, 8, 8, 9, 9} //7
	fmt.Println(Median(data))

	fmt.Println()

	dataGenap := []int{1, 2, 8, 11, 6, 10, 12, 16} // 5
	fmt.Println(Median(dataGenap))
}
