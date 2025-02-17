package main

func Receive(ch chan int) int {
	val := <- ch
	return val
}