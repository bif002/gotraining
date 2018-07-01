package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func main() {

	switch "Igor" {
	case "Dan":
		fmt.Println("Hello Dan")
	case "Sasha", "Shula":
		fmt.Println("Hello Sasha or Shula")
	case "Igor":
		fmt.Println("Hello Igor")
		fallthrough // next statement become true
	case "Idan":
		fmt.Println("Hello Idan")
	default:
		fmt.Println("No helo")
	}
	SwitchOnType(7)
	c := Contact{"hello", "Igor"}
	SwitchOnType(c)
}

// SwitchOnType : function which evaluates type of passed variable
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
		fmt.Printf("Contact : %v \n", x)
		fmt.Printf("Contact : %v \n", &x)
	default:
		fmt.Println("unknown type")
		fmt.Println(x)
		fmt.Printf("type %T \n", x)
		fmt.Printf("Default format + field names %+v \n", x)
		fmt.Printf("Representing value %#v \n", x)
		fmt.Printf("Default HEX address %v \n", &x)
	}
}
