package main

import "fmt"

func isLatin(input string) bool {
	for i := range len(input) {
		inpt := []rune(input)[i]
		fmt.Println(inpt)
		if inpt > 122 || inpt < 65 || inpt == 91 || inpt == 92 || inpt == 93 || inpt == 94 || inpt == 95 || inpt == 96 {
			return false
		}

	}
	return true
}

func main() {
	input := ("Nihao")
	fmt.Println(isLatin(input))
}
