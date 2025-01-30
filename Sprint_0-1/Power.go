/*package main

import "fmt"

func IsPowerOfTwoRecursive(N int) {

	if N < 0 {
		N *= -1
	}
	if N%2 != 0 {
		fmt.Println("No")
	} else if N/2 == 1 {
		fmt.Println("Yes")
	} else {
		IsPowerOfTwoRecursive(N / 2)
	}

}*/
package main

import "fmt"

func FindMinMaxInArray(array [10]int) (int, int) {
	a [10]int := sort(array)
	return a[1] ,1

}

