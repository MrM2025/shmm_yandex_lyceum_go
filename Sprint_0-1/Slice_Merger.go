package main

import "fmt"

func Join(nums1, nums2 []int) []int {
	length := len(nums1) + len(nums2)
	length1 := len(nums1)
	//length2 := len(nums2)

	merged_slice := make([]int, length)

	for i := 0; i < length1; i++ {
		merged_slice[i] = nums1[i]

	}

	for j, x := length1, 0; j < length; /* || x < length2 */ j++ {
		merged_slice[j] = nums2[x]
		x++

	}
	return merged_slice
}
func main() {
	input1 := []int{1, 2, 3, 4, 5, 6}
	input2 := []int{7, 8, 9, 10}
	fmt.Println(Join(input1, input2))
}
