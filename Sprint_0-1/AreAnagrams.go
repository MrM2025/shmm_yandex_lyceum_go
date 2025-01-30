package main

import (
	"fmt"
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	var str1_slice []string
	var str2_slice []string

	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")

	str1 = strings.ToUpper(str1)
	str2 = strings.ToUpper(str2)

	for _, valueofstr1 := range str1 {
		str1_slice = append(str1_slice, string(valueofstr1))
	}

	for _, valueofstr2 := range str2 {
		str2_slice = append(str2_slice, string(valueofstr2))
	}

	sort.Strings(str1_slice)
	sort.Strings(str2_slice)

	for index, _ := range str1_slice {
		if str1_slice[index] != str2_slice[index] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(AreAnagrams("Кабан", "банка"))
}
