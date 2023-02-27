package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	rand.Seed(7)
	var noOfGoroutines int
	flag.IntVar(&noOfGoroutines, "count", 0, "# of goroutines")
	flag.Parse()

	wg := &sync.WaitGroup{}
	fmt.Printf("Starting %d goroutines... HIT ENTER to start\n", noOfGoroutines)
	fmt.Scanln()
	for i := 1; i <= noOfGoroutines; i++ {
		wg.Add(1)    // increment the counter by 1
		go fn(wg, i) //scheduling the execution of this function
	}

	wg.Wait() // block until the wg counter becomes 0
	fmt.Println("main completed")
	fmt.Scanln()
}

func fn(wg *sync.WaitGroup, id int) {
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
	wg.Done() // decrement the counter by 1
}
