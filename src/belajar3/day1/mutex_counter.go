package day1

import "sync"

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.v
}

func Mutex() {
	// c := Safe
}
