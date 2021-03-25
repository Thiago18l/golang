package main

import "sync"

type Contador struct {
	mu    sync.Mutex
	value int
}

func NewCount() *Contador {
	return &Contador{}
}

func (c *Contador) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Contador) Value() int {
	return c.value
}
