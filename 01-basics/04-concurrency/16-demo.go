package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println(<-ch1)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println(<-ch2)
		wg.Done()
	}()
	wg.Wait()
}
