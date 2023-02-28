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
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")
	cancelCtx, cancel := context.WithCancel(valCtx)

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
	fmt.Println("[fn] Value from context [key=root-key]:", ctx.Value("root-key"))

	overCtx := context.WithValue(ctx, "root-key", "fn-override-value")
	valCtx := context.WithValue(overCtx, "fn-key", "fn-value")

	wg.Add(1)
	ctx2, cancel2 := context.WithTimeout(valCtx, 5*time.Second)
	defer cancel2()
	go f1(wg, ctx2)

	wg.Add(1)
	ctx3, cancel3 := context.WithTimeout(valCtx, 10*time.Second)
	defer cancel3()
	go f2(wg, ctx3)
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
	fmt.Println("[f1] Value from context [key=root-key]:", ctx.Value("root-key"))
	fmt.Println("[f1] Value from context [key=fn-key]:", ctx.Value("fn-key"))
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
	fmt.Println("[f2] Value from context [key=root-key]:", ctx.Value("root-key"))
	fmt.Println("[f2] Value from context [key=fn-key]:", ctx.Value("fn-key"))
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
