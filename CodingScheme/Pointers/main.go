package main

import (
	"fmt"
)

func main() {

	a := 10

	fmt.Println(a)  // 10
	fmt.Println(&a) // 0xc04200e0b0 - memory address

	var b = &a // var b *int = &a - linter error for infered reference

	fmt.Println(b)  // 0xc04200e0b0 - memory address of a
	fmt.Println(*b) // 10 - value of that memory address

	*b = 40 // assign new value into address storred in b

	fmt.Println(a) // 40 - new value of a (indirect assignment)
}
