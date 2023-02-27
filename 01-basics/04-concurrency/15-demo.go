package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for val := range ch {
		fmt.Println(val)
	}
}

//producer
func genNos(ch chan int) {
	for i := 1; i <= 20; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
	close(ch)
}
