package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
)

func convert(input <-chan uint8) <-chan float64 {
	output := make(chan float64)

	go func() {
		defer close(output)
		for value := range input {
			output <- math.Pow(float64(value), 3)
		}
	}()

	return output
}

func main() {

	input := make(chan uint8)
	output := convert(input)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		res := make([]float64, 0)
		for value := range output {
			res = append(res, value)
		}

		for _, value := range res { //to print firstly all input and only then - all converted values
			fmt.Printf("converted: %f\n", value)
		}
	}()

	for i := 0; i < 10; i++ {
		value := uint8(rand.Uint())
		fmt.Printf("input: %d\n", value)
		input <- value
	}

	close(input)
	wg.Wait()
}
