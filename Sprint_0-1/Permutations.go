package main

import (
	"fmt"
)

func factorial(n int) (int, error) {
	var factorial int = 1
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	for i := 1; i < n; i++ {
		factorial = factorial * (i + 1)
	}
	return factorial, nil
}

func Permutations(input string) []string {
	var count int
	var sliceofinput []string = []string{}
	var onelemslice []string = []string{}
	var newslice []string = []string{}

	for index, _ := range input {
		sliceofinput = append(sliceofinput, string(input[index]))
		//fmt.Println(sliceofinput)
	}

	for index, _ := range sliceofinput {

		//fmt.Println(sliceofinput)
	}

	for range input {
		count++
	}

	//fact, _ := factorial(count)
	return []string{}
}

func main() {
	Permutations("abc")
}
