package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	doneCh := make(chan struct{})

	wg.Add(1)
	go f1(doneCh, ch, wg)

	wg.Add(1)
	go f2(doneCh, ch, wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-doneCh:
				break LOOP
			case ch <- i * 2:
				time.Sleep(500 * time.Millisecond)
			}
		}
		fmt.Println("goroutine feeding data completed")
	}()

	go func() {
		fmt.Println("Hit ENTER to shutdown....")
		fmt.Scanln()
		close(doneCh)
		// doneCh <- struct{}{}
	}()

	wg.Wait()
	fmt.Println("Done")
}

func f1(doneCh chan struct{}, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-doneCh:
			break LOOP
		case data := <-ch:
			fmt.Println("f1 - data :", data)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("f1 completed")
}

func f2(doneCh chan struct{}, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-doneCh:
			break LOOP
		case data := <-ch:
			fmt.Println("f2 - data :", data)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("f2 completed")
}
