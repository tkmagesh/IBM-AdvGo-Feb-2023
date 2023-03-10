package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg.Add(1)
	go f1(cancelCtx, ch, wg)

	wg.Add(1)
	go f2(cancelCtx, ch, wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-cancelCtx.Done():
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
		cancel()
	}()

	wg.Wait()
	fmt.Println("Done")
}

func f1(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case data := <-ch:
			fmt.Println("f1 - data :", data)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("f1 completed")
}

func f2(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case data := <-ch:
			fmt.Println("f2 - data :", data)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("f2 completed")
}
