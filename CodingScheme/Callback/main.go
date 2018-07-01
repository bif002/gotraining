package main

import "fmt"

func visit(numbers []int, clb func(int)) { // visit function accepting 2 params where the last one is the func
	for _, n := range numbers {
		clb(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) { // definition of callback function
		fmt.Println(n)
	})
}
