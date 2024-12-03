package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func sendValue[T comparable](channel chan T, value T, sentValues *sync.Map) {
	channel <- value
	sentValues.Store(value, true)
}

func TestMerge(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	var sentValues sync.Map

	go func(ch chan int) {
		sendValue(ch, 10, &sentValues)
		sendValue(ch, 20, &sentValues)
		sendValue(ch, 30, &sentValues)
		close(ch)
	}(ch1)
	go func(ch chan int) {
		sendValue(ch, 40, &sentValues)
		sendValue(ch, 50, &sentValues)
		sendValue(ch, 60, &sentValues)
		close(ch)
	}(ch2)
	go func(ch chan int) {
		sendValue(ch, 70, &sentValues)
		sendValue(ch, 80, &sentValues)
		sendValue(ch, 90, &sentValues)
		close(ch)
	}(ch3)

	merged := merge(ch1, ch2, ch3)
	received := make([]int, 0)

	for value := range merged {
		received = append(received, value)
	}

	for _, value := range received {
		_, ok := sentValues.Load(value)
		assert.True(t, ok)
		sentValues.Delete(value)
	}

	empty := true
	sentValues.Range(func(key, value any) bool {
		empty = false
		return false
	})

	assert.True(t, empty)
}
