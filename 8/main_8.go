package main

import (
	"fmt"
	"sync"
	"time"
)

type CustomWaitGroup struct {
	cond    *sync.Cond
	mu      *sync.Mutex
	counter int
}

func newCustomWaitGroup() *CustomWaitGroup {
	mu := sync.Mutex{}
	return &CustomWaitGroup{cond: sync.NewCond(&mu), mu: &mu}
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.counter += delta
	if delta < 0 {
		wg.checkCounter()
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.counter--
	wg.checkCounter()
}

func (wg *CustomWaitGroup) Wait() {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	for wg.counter > 0 {
		wg.cond.Wait()
	}
}

func (wg *CustomWaitGroup) checkCounter() {
	if wg.counter == 0 {
		wg.cond.Broadcast()
	} else if wg.counter < 0 {
		panic("wait group counter is negative")
	}
}

func main() {

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
