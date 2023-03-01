package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go f1(ch)
	go f2(ch)
	go func() {
		for i := 1; ; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 2
		}
	}()
	fmt.Scanln()
}

func f1(ch chan int) {
	for data := range ch {
		time.Sleep(1 * time.Second)
		fmt.Println("f1 - data :", data)
	}
}

func f2(ch chan int) {
	for data := range ch {
		time.Sleep(2 * time.Second)
		fmt.Println("f2 - data :", data)
	}
}
