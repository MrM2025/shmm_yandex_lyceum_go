package main

import "fmt"

func Mix(nums []int) []int {
	length := len(nums)
	if length%2 != 0 {
		return nil
	}
	middle := length / 2
	for x, y := 0, middle; y < length; x++ {
		q := nums[x+1]
		nums[x+1] = nums[y]
		nums[x+2] = q

		y++
	}
	return nums

}

func main() {
	input := []int{1, 2, 3, 4, 5, 11, 12, 13, 14, 15}
	fmt.Println(Mix(input))
}
