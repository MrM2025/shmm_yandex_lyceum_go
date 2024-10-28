package main

import "fmt"

func IsPalindrome(input string) bool {
	for i := range len(input) / 2 {
		length := len(input)
		length--
		fmt.Println(input[0])

		if input[i] != input[length] {
			fmt.Println(input[i], " ", input[length])
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("A B C C B A"))
}
