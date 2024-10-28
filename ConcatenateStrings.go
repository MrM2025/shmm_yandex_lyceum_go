package main

import "fmt"

func ConcatenateStrings(str1, str2 string) string {
	ConcStr := str1 + " " + str2
	return ConcStr
}

func main() {
	str1 := ("Anna")
	str2 := ("Svistakova")
	fmt.Println(str1, str2)
}
