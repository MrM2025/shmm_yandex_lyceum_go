package main

import (
	"fmt"
	"strconv"
)

func SumTwoIntegers(a, b string) (int, error) {
	num1, err := strconv.Atoi(a)
	num2, err1 := strconv.Atoi(b)
	if err != nil || err1 != nil {
		return 0, fmt.Errorf("invalid input, please provide two integers")
	}
	return num1 + num2, nil

}

func main() {
	fmt.Println(SumTwoIntegers("1", "2"))
}
