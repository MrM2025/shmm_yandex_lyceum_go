package main

import (
	"sync"
)

type Queue interface {
	Enqueue(element interface{}) // положить элемент в очередь
	Dequeue() interface{} // забрать первый элемент из очереди
	}

type ConcurrentQueue struct {
	queue []interface{}	 // здесь хранить элементы очереди
	mutex sync.Mutex
	}

func (c *ConcurrentQueue) Enqueue(element interface{}) {
	
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.queue = c.queue[1:]
	c.queue = append(c.queue, element)
}

func (c *ConcurrentQueue) Dequeue() interface{} {
	return c.queue[0]
}
