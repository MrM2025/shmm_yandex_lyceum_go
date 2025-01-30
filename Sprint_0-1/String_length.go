package main

import "fmt"

func StringLength(input string) int {
	var length int
	if input == "" {
		length = 0
	}
	length = len(input)
	return length

}

func main() {
	input := ("eyjafjallajkull")
	fmt.Println(StringLength(input))
}
