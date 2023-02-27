package main

import "fmt"

type OperationFn func(int, int) int

func main() {
	/*
		fn := func() {
			fmt.Println("fn invoked")
		}
	*/

	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	// var add func(int, int) int
	var add OperationFn
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	// var subtract func(int, int) int
	var subtract OperationFn
	subtract = func(i1, i2 int) int {
		return i1 - i2
	}
	fmt.Println(subtract(100, 200))
}
