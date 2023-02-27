package main

import "fmt"

func main() {
	anonFn := getFn()
	anonFn()
}

func getFn() func() {
	return func() {
		fmt.Println("anon fn invoked")
	}
}
