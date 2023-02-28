package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	go func() {
		fmt.Println("Hit ENTER to stop")
		fmt.Scanln()
		cancel()
	}()
	wg.Add(1)
	go fn(wg, cancelCtx)
	wg.Wait()
}

func fn(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	wg.Add(1)
	go f1(wg, ctx)

	wg.Add(1)
	go f2(wg, ctx)
LOOP:
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("[fn] Cancel signal received")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("fn :", i*10)
		}
	}
}

func f1(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
LOOP:
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("[f1] Cancel signal received")
			break LOOP
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("f1 :", i*2)
		}
	}
}

func f2(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
LOOP:
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("[f2] Cancel signal received")
			break LOOP
		default:
			time.Sleep(700 * time.Millisecond)
			fmt.Println("f2 :", (i*2)+1)
		}
	}
}
