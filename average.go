package main

import "fmt"

func Average(value []float64) float64 {
	lengthVal := float64(len(value))
	var totalVal float64 = 0

	for _, e := range value {
		totalVal += e
	}

	avg := totalVal / lengthVal
	return avg
}

func main() {
	data := []float64{1, 2, 3, 4, 5}
	fmt.Println(Average(data))
}
