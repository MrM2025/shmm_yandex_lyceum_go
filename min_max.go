package main

func FindMinMaxInArray(array [10]int) (int, int) {

	var a int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if array[i] < array[j] {
				a = array[i]
				array[i] = array[j]
				array[j] = a
			}
		}
	}
	return array[0], array[9]
}
