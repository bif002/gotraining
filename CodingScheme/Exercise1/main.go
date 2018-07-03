package main

import "fmt"

// Run program that accept int value and returns
// 2 arguments
//   1-st - accepted value devided by 2
//   2-nd - is boolean where the parameter even or not
func half(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}

func main() {
	n, even := half(5)
	fmt.Println(n, even)
}
