package main

import "fmt"

func main() {
	// var x interface{}
	var x any
	// x = 100
	// x = "This is a string"
	// x = true
	// x = 99.99
	// x = struct{}{}
	fmt.Println(x)

	//type assertion using if
	x = 200
	// x = true
	// fmt.Println(x.(int) + 300)
	if val, ok := x.(int); ok {
		fmt.Println(val + 300)
	} else {
		fmt.Println("x is not an int")
	}

	//type assertion using switch
	// x = true
	x = "Veniam cillum laboris laboris esse commodo quis sit. Non ea elit sunt sit labore ea minim Lorem. Velit velit velit quis occaecat magna ut."
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, 99.9% of x is", val*(99.9/100))
	default:
		fmt.Println("Unknown type")
	}
}
