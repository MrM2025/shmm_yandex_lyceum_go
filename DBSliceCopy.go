package main

import (
	"fmt"
)

func SliceCopy(nums []int) []int {
	length := len(nums)
	slice := make([]int, length) //?????????????? why without length don't work
	for i := 0; i < length; i++ {
		slice[i] = nums[i]
	}

	return slice
}

func main() {
	input := make([]int, 0, 10)

	fmt.Println(SliceCopy(input))
}
