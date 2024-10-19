package main

import "fmt"

func FibonacciRecursivePrint(iterator, previous_digit, current_digit int) int {
	if iterator < 3 {
		return -1
	} else if previous_digit == 0 {
		fmt.Println(previous_digit, current_digit)
	}
	fmt.Println(previous_digit + current_digit)
	return FibonacciRecursivePrint(iterator-1, current_digit, previous_digit+current_digit)
}

func main() {

	var iterator int
	fmt.Println("Введите длину ряда чисел Фибоначчи")
	fmt.Scanln(&iterator)
	if iterator < 3 {
		fmt.Println("Неверная длина! Значение длины ряда должно быть больше двух")
	} else {
		fmt.Println("Считаем и выводим рекурсивно")
		FibonacciRecursivePrint(iterator, 0, 1)
		fmt.Println("Считаем и выводим в цикле")
		previous_digit, current_digit := 0, 1
		fmt.Println(previous_digit)
		for iterator > 1 {
			fmt.Println(current_digit)
			previous_digit, current_digit = current_digit, previous_digit+current_digit
			iterator--
		}
	}

}
