package main

import (
	"errors"
	"fmt"
)

func SearchCharacter(str string, value string) (bool, error) {
	if len(str) == 0 {
		return false, errors.New("Empty string")
	}

	var search bool

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == value[0] {
			search = true
			break
		}
	}

	return search, nil
}

func main() {
	result, err := SearchCharacter("Aku Adalah Raja", "A")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(result)

	result2, err2 := SearchCharacter("Gtw OhA", "A")
	if err2 != nil {
		fmt.Println("Error:", err2)
	}

	fmt.Println(result2)
}
