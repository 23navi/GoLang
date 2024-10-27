package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	doubleNumbers(numbers)
	fmt.Println(numbers) // [1 2 3 4]
}

func doubleNumbers(numbers []int) {
	// dNumbers := []int{}

	for i, val := range numbers {
		// dNumbers = append(dNumbers, val*2)
		numbers[i] = val * 2
	}
}
