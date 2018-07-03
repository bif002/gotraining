package main

import "fmt"

// Write a Variadic function which gets a list of numbers
// and returns the larges number in the list

func max(numbers ...int) int {
	// fmt.Printf("%T\n", numbers)
	var largets int
	for _, v := range numbers {
		// fmt.Println(k, v)
		if v > largets {
			largets = v
		}
	}
	return largets
}

func main() {
	largest := max(2, 4, 3, 34, 65, 7, 93)
	fmt.Println(largest)
}
