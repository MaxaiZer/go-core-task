package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	wg := newCustomWaitGroup()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(i) * 200 * time.Millisecond)
			fmt.Printf("%d finished\n", i)
		}()
	}

	fmt.Println("waiting for goroutines...")
	wg.Wait()
	fmt.Println("all goroutines finished")
}

func TestWaitGroup_Panic(t *testing.T) {
	assert.Panics(t, func() {
		wg := newCustomWaitGroup()
		wg.Add(-1)
	})
}
