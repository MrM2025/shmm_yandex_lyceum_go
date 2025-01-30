package main

import (
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	length := len(nums)
	if n <= 0 || nums == nil || limit == 0 {
		return nums, fmt.Errorf("wrong, %v", n)
	}

	ns_under_limit := []int{}
	for i := 0; i < length; i++ {

		if nums[i] < limit {
			ns_under_limit = append(ns_under_limit, nums[i])
		}
	}
	return ns_under_limit, nil
}

func main() {
	//a := []int{3, 5, 6}
	fmt.Println(UnderLimit([]int{}, 5, 2))

}
