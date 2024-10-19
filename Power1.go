package main

import "fmt"

func IsPowerOfTwoRecursive(N int) {

	if N == 1 {
		fmt.Println("YES")
		return
	}

	if N%2 != 0 {
		fmt.Println("NO")
		return
	} else if N/2 == 1 {
		fmt.Println("YES")
		return
	} else {
		IsPowerOfTwoRecursive(N / 2)
	}

	return

}
