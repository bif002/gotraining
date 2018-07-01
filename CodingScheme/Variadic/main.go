package main

import "fmt"

func main() {
	n := average(23, 45, 67, 54, 12, 1)
	fmt.Println(n)
	data := []float64{23, 45, 67, 54, 12, 1}
	// calling function average and passing variadic parameters from slice of []float64
	// float64 and []float64 are completely different types
	// ... says that instead of passing one peace / slice from the list, pass all values
	// like in previous example
	y := average(data...)
	fmt.Println("Passing variadic arguments ", y)

	z := average2(data)
	fmt.Println("Passing the same type", z)
}

// Defining function which accept unlimitted number of arguments
// AKA : Variadic parameters to function
func average(x ...float64) float64 {
	var total float64
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}

func average2(x []float64) float64 {
	var total float64
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}
