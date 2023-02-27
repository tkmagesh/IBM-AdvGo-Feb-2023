package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		stopCh <- struct{}{}
	}()
	ch := generateNos(stopCh)
	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("Done")
}

func generateNos(stopCh chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-stopCh:
				break LOOP
			default:
				ch <- i * 2
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}
