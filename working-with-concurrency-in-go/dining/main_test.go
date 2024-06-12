package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 1 * time.Second
	thinkTime = 3 * time.Second
	sleepTime = 1 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()

		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slices; expected 5 but got %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second dealy", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		orderFinished = []string{}

		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s: incorrect length of slice; expected 5 but got %d", e.name, len(orderFinished))
		}
	}
}
