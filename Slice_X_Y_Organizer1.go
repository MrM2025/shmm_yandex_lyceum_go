package main

import "fmt"

func Mix(nums []int) []int {
	length := len(nums)
	middle := length / 2

	slice := []int{}

	for x, y := 0, middle; y < length; x++ {
		slice = append(slice, nums[x])
		slice = append(slice, nums[y])
		y++
	}
	return slice

}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Mix(input))

}
