package main

import "fmt"

func Factorial(n int) (int, error) {
	var factorial int = 1
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	for i := 1; i < n; i++ {
		factorial = factorial * (i + 1)
	}
	return factorial, nil
}

func main() {
	fmt.Println(Factorial(6))
}
