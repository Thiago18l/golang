package main

func newStoragePlayerInMemory() *PlayerStorageInMemory {
	return &PlayerStorageInMemory{map[string]int{}}
}

type PlayerStorageInMemory struct {
	storage map[string]int
}

func (p *PlayerStorageInMemory) RegisterOwn(name string) {
	p.storage[name]++
}

func (p *PlayerStorageInMemory) GetPlayerPoints(name string) int {
	return p.storage[name]
}
