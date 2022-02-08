package main

import "fmt"

func Factorial(value int) int {
	if value <= 0 {
		return 1
	}

	result := 1
	for i := value; i >= 1; i-- {
		result *= i
	}

	return result
} //Solved

func FactorialRecursive(value int) int {
	if value <= 0 {
		return 1
	} else {
		return value * Factorial(value-1)
	}
} //Solved

func FactorialRecursiveTail(total, value int) int {
	if value <= 0 {
		return total
	} else {
		return FactorialRecursiveTail(total*value, value-1)
	}
} //Solved

func main() {
	fmt.Println(Factorial(5))
	fmt.Println(FactorialRecursive(5))
	fmt.Println(FactorialRecursiveTail(1, 5))
	fmt.Println()
	fmt.Println()
	fmt.Println(Factorial(3))
	fmt.Println(FactorialRecursive(3))
	fmt.Println(FactorialRecursiveTail(1, 3))
}
