package main

import "fmt"

func FindMaxKey(m map[int]int) int {
	var max int
	// How can I assign  one key to the variable
	//max = key
	fmt.Println(key)
	for key1, _ := range m {

		fmt.Println(key1)
		if max > key1 {
		}
		if max < key1 {
			max = key1
		}
	}
	return max
}

func main() {
	map1 := map[int]int{71: 1, -7: 0, -3217: 0, 1237: 0, 127: 0, 7: 0}
	fmt.Println(FindMaxKey(map1))
}
