package main

import "fmt"

func zero2(z *int) {
	fmt.Println(z)
	*z = 0
}

func zero(x int) {
	fmt.Printf("Inside zero : %p", &x)
	x = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // still printing 5 coz passing func args always by value
	fmt.Println(&x)
	zero2(&x)      // call func zero2 with address reference to x
	fmt.Println(x) // 0
}
