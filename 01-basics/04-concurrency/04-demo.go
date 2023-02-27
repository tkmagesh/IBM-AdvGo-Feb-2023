package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Counter :", counter)
}

func increment(wg *sync.WaitGroup) {
	counter++
	wg.Done()
}
