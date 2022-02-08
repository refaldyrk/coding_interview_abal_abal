package main

import (
	"fmt"
)

func FizzBuzz(value int) {
	if value <= 0 {
		fmt.Println("No Zero")
	}

	for i := 1; i <= value; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if 1%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	FizzBuzz(100)
}
