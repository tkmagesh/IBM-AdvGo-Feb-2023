package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // increment the counter by 1
	go f1()   //scheduling the execution of this function
	f2()
	wg.Wait() // block until the wg counter becomes 0
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
