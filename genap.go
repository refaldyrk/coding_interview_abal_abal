package main

import "fmt"

func FindGenapSingle(value int) bool {
	if value%2 == 0 {
		return true
	}

	return false
}

/*
Cari Nilai Angka Genap Di Dalam Array Or Slice
Dan Function Ini Akan Mengembalikan Array Of Slice Yang Bertipe Integer
*/
func FindGenapArray(value []int) []int {
	var newArr []int
	for _, e := range value {
		if e%2 == 0 {

			newArr = append(newArr, e)
		}
	}

	return newArr
}

func main() {
	fmt.Println(FindGenapSingle(10)) // true
	fmt.Println(FindGenapSingle(9))  // false

	data := []int{1, 2, 3, 4, 5} // 2 4
	fmt.Println(FindGenapArray(data))
}
