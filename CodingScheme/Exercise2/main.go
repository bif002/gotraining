package main

import "fmt"

// Modify previous Ex1 program to use
// func expression

func main() {
	half := func(x int) (float64, bool) {
		return float64(x) / 2, x%2 == 0
	}

	n, even := half(5)
	fmt.Println(n, even)
}
