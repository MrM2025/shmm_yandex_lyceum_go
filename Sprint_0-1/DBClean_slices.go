package main

import (
	"fmt"
)

func Clean(nums []int, x int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			nums = append(nums[:i], nums[i+1:]...) // ??? (...)
			i--
		}
	}
	return nums
}

func main() {
	input := []int{3, 20, 3, 4, 5, 1, 3, 7, 9, 12}
	fmt.Println(Clean(input, 3))

}
