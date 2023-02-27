package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	dataCh := make(chan int)

	wg.Add(1)
	go source("data1.dat", dataCh, wg)

	wg.Add(1)
	go source("data2.dat", dataCh, wg)

	evenCh, oddCh := splitter(dataCh)
	evenSumCh := sum(evenCh)
	oddSumCh := sum(oddCh)

	doneCh := merger(evenSumCh, oddSumCh, "result.dat")
	wg.Wait()
	close(dataCh)

	<-doneCh
	fmt.Println("Done!")
}

func source(fileName string, ch chan int, wg *sync.WaitGroup) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if val, err := strconv.Atoi(txt); err == nil {
			ch <- val
		}
	}
	wg.Done()
}

func splitter(ch chan int) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)
	go func() {
		for val := range ch {
			if val%2 == 0 {
				evenCh <- val
			} else {
				oddCh <- val
			}
		}
		close(evenCh)
		close(oddCh)
	}()
	return evenCh, oddCh

}

func sum(valCh <-chan int) <-chan int {
	resultCh := make(chan int)
	go func() {
		result := 0
		for val := range valCh {
			result += val
		}
		resultCh <- result
	}()
	return resultCh
}

func merger(evenSumCh, oddSumCh <-chan int, fileName string) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		for i := 0; i < 2; i++ {
			select {
			case evenSum := <-evenSumCh:
				fmt.Fprintf(file, "Even Total : %d\n", evenSum)
			case oddSum := <-oddSumCh:
				fmt.Fprintf(file, "Odd Total : %d\n", oddSum)
			}
		}
		doneCh <- struct{}{}
	}()
	return doneCh
}
