package main

import (
	"fmt"
	"math/rand"
)

// функция генерирует случайное число в интервале [0, 100)
func random() int {
	const max int = 100
	return rand.Intn(max)
}

func main() {
	const size int = 10
	results := []int{}
	// заполняем слайс случайными числами
	for i := 0; i < size; i++ {
		results = append(results, random())
	}

	// поэлементно выводим слайс на экран
	for i := 0; i < size; i++ {
		fmt.Println(results[i])
	}
}