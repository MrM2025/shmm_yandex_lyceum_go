package main

import "fmt"

func PopNum(sliceofnums []float64, numtopop int) ([]float64, []float64, error) {

	var popednum, newsliceofnums []float64

	if numtopop > len(sliceofnums) {
		return popednum, sliceofnums, fmt.Errorf("index of num > length of slice of nums, %d", numtopop)
	}
	if numtopop <= 0 {
		return popednum, sliceofnums, fmt.Errorf("index of num <= 0, %d", numtopop)
	}

	for countdownindex := 1; countdownindex <= numtopop; countdownindex++ {
		popednum = append(popednum, sliceofnums[len(sliceofnums)-countdownindex])
		newsliceofnums = append(sliceofnums[:len(sliceofnums)-numtopop], sliceofnums[len(sliceofnums):]...)
		fmt.Println(newsliceofnums, popednum, len(sliceofnums))
	}

	return popednum, newsliceofnums, nil
}

func main() {
	var a, b []float64

	for i := 0; i < 10; i++ {
		a = append(a, float64(i))
	}

	fmt.Println(a)
	b, a, _ = PopNum(a, 3)
	fmt.Println(b, a)

}
