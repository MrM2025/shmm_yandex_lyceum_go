package main

import "fmt"

func GetCharacterAtPosition(str string, position int) (rune, error) {
	for i, val := range str {
		if position > i {
			return 0, fmt.Errorf("position out of range")
		}
		if position == i {
			return val, nil
		}
	}
	return 0, nil
}

func main() {
	fmt.Println(GetCharacterAtPosition("Зенит чемпион!", 20))
}
