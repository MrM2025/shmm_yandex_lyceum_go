package main

import (
	"sync"
)

var Buf []int
var mu sync.Mutex

func Write(num int) {
	mu.Lock()
	defer mu.Unlock()
	Buf = append(Buf, num)
}

func Consume() int {
	mu.Lock()
	defer mu.Unlock()
	if len(Buf) == 0 {
		return 0
	}
	val := Buf[0]
	Buf = Buf[1:]
	return val
}