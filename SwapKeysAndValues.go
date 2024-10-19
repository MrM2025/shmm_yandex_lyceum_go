package main

import (
	"fmt"
)

func SwapKeysAndValues(m map[string]string) map[string]string {
	map1 := map[string]string{}
	for key, value := range m {
		map1[value] = key
		fmt.Println(m, key, value)
	}
	return map1

}

func main() {
	map2 := map[string]string{"1": "Hi", "2": "John", "3": "!"}
	fmt.Println(SwapKeysAndValues(map2))
}
