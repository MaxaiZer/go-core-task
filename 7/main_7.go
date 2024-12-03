package main

import (
	"fmt"
	"sync"
)

func merge[T any](channel ...chan T) chan T {
	merged := make(chan T, 10)
	wg := sync.WaitGroup{}

	for _, ch := range channel {
		wg.Add(1)
		go func() {
			for value := range ch {
				merged <- value
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func(ch chan int) {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}(ch1)
	go func(ch chan int) {
		ch <- 4
		ch <- 5
		ch <- 6
		close(ch)
	}(ch2)
	go func(ch chan int) {
		ch <- 7
		ch <- 8
		ch <- 9
		close(ch)
	}(ch3)

	merged := merge(ch1, ch2, ch3)
	for value := range merged {
		fmt.Printf("got value from merged channel: %v\n", value)
	}
}
