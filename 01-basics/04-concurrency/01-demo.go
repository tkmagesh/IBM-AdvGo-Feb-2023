package main

import (
	"fmt"
)

func main() {
	go f1() //scheduling the execution of this function
	f2()

	/* DO NOT DO THE FOLLOWING */
	// time.Sleep(1 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
