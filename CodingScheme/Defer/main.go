package main

import "fmt"

func world() {
	fmt.Println("world")
}

func hello() {
	fmt.Print("hello ")
}

func main() {
	defer world() // defer keyword pospond calling to specific statement untill funcion it was called in
	// exiting (define the last execution in specific block / function)
	hello()
}
