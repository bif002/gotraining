package main

import (
	"fmt"
)

const (
	/* The following comments are MANDATORY for linter to not
	through warning message like
	exported const PI should have comment (or a comment on this block) or be unexported */

	// PI : define constant Pi
	// PI = 3.14
	// Language : define language const
	// Language = "Go"

	/*
		In order to get proper resuts with iota, the iota declaration have to be first and
		the only declaration in constant block
	*/

	_ = iota // 0
	// KB : iota is running tiny autoincrementor
	KB = 1 << (iota * 10) //1 << (1 * 10)
	// MB : next call to iota couse to increment by 1
	MB = 1 << (iota * 10) //1 << (2 * 10)
)

func main() {

	const q = 42
	// fmt.Println(PI)
	// fmt.Println(Language)
	fmt.Printf("KB - %b in dec %d \n", KB, KB)
	fmt.Printf("MB - %b in dec %d \n", MB, MB)
}
