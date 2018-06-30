package main

import (
	"fmt"
)

const (
	/* The following comments are MANDATORY for linter to not
	through warning message like
	exported const PI should have comment (or a comment on this block) or be unexported */

	// PI : define constant Pi
	PI = 3.14
	// Language : define language const
	Language = "Go"
)

func main() {

	const q = 42
	fmt.Println(PI)
	fmt.Println(Language)
}
