package main

import "fmt"

func CountingSort(contacts []string) map[string]int {
	map1 := map[string]int{}

	for _, value := range contacts {
		map1[value]++
	}
	return map1
}

func main() {
	input := []string{"1", "2", "3", "4", "5", "6", "1", "7", "4", "5"}
	fmt.Println(CountingSort(input))
}
