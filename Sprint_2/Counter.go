package main

import (
	"sync"
)

type Count interface{
	Increment() // увеличение счётчика на единицу
	GetValue() int // получение текущего значения
	}

type Counter struct {
	value int // значение счетчика
	mu sync.RWMutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value += 1
	c.mu.Unlock()
}

func (c Counter) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
	
}	