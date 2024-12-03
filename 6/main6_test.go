package main

import (
	"testing"
	"time"
)

func TestRand(t *testing.T) {

	randChan := customRand()

	for i := 0; i < 100; i++ {
		select {
		case _ = <-randChan:
			//can't test uniqueness because technically values can be repeated
		case <-time.After(1 * time.Second):
			t.Error("time out")
		}
	}
}
