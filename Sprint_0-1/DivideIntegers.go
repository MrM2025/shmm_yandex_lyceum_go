package main

import "fmt"

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	c := float64(a / b)
	return (c), nil
}

func main() {
	fmt.Println(DivideIntegers(1, 0))
}
