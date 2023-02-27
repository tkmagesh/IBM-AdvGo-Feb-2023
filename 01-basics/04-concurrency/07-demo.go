package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

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
	atomic.AddInt64(&counter, 1)
	wg.Done()
}
