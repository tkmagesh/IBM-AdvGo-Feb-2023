package main

import (
	"fmt"
	"log"
	"time"
)

type OperationFn func(int, int)

func main() {
	// log
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	// profile
	/*
		profileAdd := getProfileOperation(add)
		profileAdd(100, 200)

		profileSubtract := getProfileOperation(subtract)
		profileSubtract(100, 200)
	*/

	// log + profile
	logAdd := getLogOperation(add)
	profileLogAdd := getProfileOperation(logAdd)
	profileLogAdd(100, 200)
}

func getProfileOperation(operation OperationFn) OperationFn {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elpased :", elapsed)
	}
}

func getLogOperation(operation OperationFn) OperationFn {
	return func(x, y int) {
		log.Println("Operation started")
		operation(x, y)
		log.Println("Operation completed")
	}
}

//3rd party library (CANNOT change)
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
