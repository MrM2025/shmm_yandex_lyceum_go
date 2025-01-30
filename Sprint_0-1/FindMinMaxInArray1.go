package main

func FindMinMaxInArray(array [10]int) (int, int) {
	for i := 0; i < 10; i++ {
		for q := 0; q < 10; q++ {

			if array[i] > array[q] {

				a := array[q]
				array[q] = array[i]
				array[i] = a
			}

		}

	}
	return array[0], array[9]

}
