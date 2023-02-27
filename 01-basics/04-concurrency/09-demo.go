package main

import (
	"fmt"
	"sync"
)

//share memeory by communicating

func main() {
	var ch chan int
	ch = make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, 100, 200, ch)
	result := <-ch
	wg.Wait()
	fmt.Println(result)
}

func add(wg *sync.WaitGroup, x, y int, ch chan int) {
	defer wg.Done()
	result := x + y
	ch <- result
}
