package main

import "fmt"

func PopNum(sliceofnums []float64, numtopop int) ([]float64, []float64, error) {

	var newsliceofnums []float64

	if numtopop > len(sliceofnums) {
		return nil, sliceofnums, fmt.Errorf("index of num > length of slice of nums, %d", numtopop)
	}
	if numtopop <= 0 {
		return nil, sliceofnums, fmt.Errorf("index of num <= 0, %d", numtopop)
	}

	popednum := make([]float64, numtopop)
	popednum = append(sliceofnums[len(sliceofnums)-numtopop:])
	newsliceofnums = append(sliceofnums[:len(sliceofnums)-numtopop], sliceofnums[len(sliceofnums):]...)
	return popednum, newsliceofnums, nil
}

func PopOp(opslice []int) (int, []int, error) {

	var newopslice []int

	if len(opslice) == 0 {
		return 0, opslice, fmt.Errorf("no operator to pop")
	}

	popedop := opslice[len(opslice)-1]
	newopslice = append(opslice[:len(opslice)-1], opslice[len(opslice):]...)

	return popedop, newopslice, nil
}

func main() {
	var a, b []float64
	var d int

	for i := 0; i < 10; i++ {
		a = append(a, float64(i))
	}

	fmt.Println(a)
	b, a, _ = PopNum(a, 3)
	fmt.Println(b, a)

	c := []int{10, 20, 30, 40}

	d, c, _ = PopOp(c)

	fmt.Println(d, c)

}
