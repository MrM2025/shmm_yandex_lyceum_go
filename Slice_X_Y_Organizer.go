package main

import "fmt"

func Mix(nums []int) []int {
	length := len(nums)
	middle := length / 2
	if length%2 != 0 {
		length -= 1
	}
	for x, y := 0, middle; y < length; x++ {
		q := nums[x+1]
		nums[x+1] = nums[y]
		nums[x+2] = q

		y++
	}
	return nums

}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Mix(input))

}
