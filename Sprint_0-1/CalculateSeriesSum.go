package main

import "fmt"

func CalculateSeriesSum(n int) float64 {

	if n > 1 {
		return 1/float64(n) + CalculateSeriesSum(n-1)
	}

	return 1
}

func main() {
	fmt.Println(CalculateSeriesSum(6))
}

/*
package main


func CalculateSeriesSum(n int) float64 {
	var c float64
	b := float64(n)

	for b > 0 {
		c += 1 / b
		b--
	}
	return c

}
*/
