package main

import "fmt"

func DeleteLongKeys(m map[string]int) map[string]int {
	map1 := make(map[string]int)
	for key, value := range m {
		if len(key) >= 6 {
			map1[key] = value
		}
	}
	return map1
}

func main() {
	map1 := map[string]int{"1234567": +1234, "123": 13, "123456": 12334, "12344566664": 12}
	fmt.Println(DeleteLongKeys(map1))
}
