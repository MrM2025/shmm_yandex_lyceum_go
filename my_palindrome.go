package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(input string) bool {
	length := len(input) - 1
	forlength := (length + 1) / 2
	j := 0
	input = strings.ToUpper(input)
	for i := range forlength {
		if string(input[i]) == " " {
			j++
			continue
		}
		for {
			if string(input[length-i+j]) != " " {
				break
			}
			j--
		}
		if input[i] != input[length-i+j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("Was it a cat I saw"))
}
