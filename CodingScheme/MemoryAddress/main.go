package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Printf("%f meters is %f yards \n", meters, yards)
	a := 42
	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
}
