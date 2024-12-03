package main

import (
	"fmt"
	"math/rand"
)

func customRand() <-chan int {

	ch := make(chan int)

	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	return ch
}

func main() {

	randChan := customRand()

	for i := 0; i < 10; i++ {
		fmt.Println(<-randChan)
	}
}
