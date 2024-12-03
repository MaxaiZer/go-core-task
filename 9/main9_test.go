package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand/v2"
	"sync"
	"testing"
)

func TestConvert(t *testing.T) {

	ch1 := make(chan uint8)
	ch2 := convert(ch1)

	inputValues := make([]uint8, 0)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		res := make([]float64, 0)
		for value := range ch2 {
			res = append(res, value)
		}

		for _, value := range res {
			assert.Equal(t, math.Pow(float64(inputValues[0]), 3), value)
			inputValues = inputValues[1:]
		}
	}()

	for i := 0; i < 10; i++ {
		value := uint8(rand.Uint())
		inputValues = append(inputValues, value)
		ch1 <- value
	}

	close(ch1)
	wg.Wait()
}
