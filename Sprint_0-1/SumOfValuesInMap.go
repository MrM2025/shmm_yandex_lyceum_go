package main

import "fmt"

func SumOfValuesInMap(m map[int]int) int {
	var sum int
	for _, values := range m {
		sum += values
	}
	return sum
}

func main() {
	map1 := map[int]int{1: 2, 2: 3, 3: 10, 4: 10, 5: 10000000}
	fmt.Println(SumOfValuesInMap(map1))
}
