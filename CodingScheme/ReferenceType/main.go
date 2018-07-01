package main

import "fmt"

// Everythin is passed by VALUE not by reference

func main() {
	m := make([]string, 1, 24) // the make keyword here is defining m variable being passed by reference/
	// or passwing reference / address of the array string
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "Igor"
	fmt.Println(z)
}
