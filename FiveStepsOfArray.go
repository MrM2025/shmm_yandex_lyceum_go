package main

import "fmt"

func FiveSteps(array [5]int) [5]int { // func Array_Reverse
	c := 4
	i := 0
	for i < 2 {

		a := array[i]
		array[i] = array[c]
		array[c] = a
		i++
		c--
		fmt.Println(a, array)

	}
	return array
}
