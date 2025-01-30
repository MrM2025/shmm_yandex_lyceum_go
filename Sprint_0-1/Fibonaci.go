package main

import "fmt"

func FibonaciSlover() int {
	var val int = 1
	for i := range 6 {
		val1 := val
		val = val + val
		fmt.Println(val, i)
	}
	return val
}

func main() {
	fmt.Println(FibonaciSlover())
}
