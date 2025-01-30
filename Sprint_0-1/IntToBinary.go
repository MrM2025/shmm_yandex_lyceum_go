package main

import "fmt"

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", fmt.Errorf("negative numbers are not allowed")
	}
	return fmt.Sprintf("%b", num), nil
}

func main() {
	fmt.Println(IntToBinary(10))
}
