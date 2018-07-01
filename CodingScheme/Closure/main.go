package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()

	fmt.Printf("The type of increment var is %T \n", increment)

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

	// The results as that because wrapper() function return function
	// where x in a outer scope and the actual function execution
	// is in the inner scope which is actually increments x by 1
}
