package main

import "fmt"

func main() {
	var current_num int
	previous_num := 0
	fib_num := 1
	fib_num += previous_num
	fmt.Scanln(&current_num)
	for i := 5; i > 0; i-- {
		fmt.Println(previous_num, fib_num)

		previous_num += fib_num
		fib_num += previous_num

	}
}

/*
package main

import "fmt"

func fibonacciRecursive(iterator, previous_digit, current_digit, after_digit int) int {
	if iterator < 3 {
		return -1
	}
	if after_digit > previous_digit {
		fmt.Println(after_digit, previous_digit, current_digit)
		iterator++

	}
	if after_digit <= previous_digit+current_digit {
		fmt.Println(previous_digit + current_digit)
	}
	return fibonacciRecursive(iterator-1, current_digit, previous_digit+current_digit, after_digit)

}

func main() {
	var num int
	fmt.Scanln(&num)
	fibonacciRecursive(10, 0, 1, num)
}
*/
