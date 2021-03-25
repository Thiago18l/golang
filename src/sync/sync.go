package main

type Contador struct {
	value int
}

func (c *Contador) Increment() {
	c.value++
}

func (c *Contador) Value() int {
	return c.value
}
