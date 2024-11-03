package main

import (
	"fmt"
)

func div(num int) int { // вычисляем делитель
	div := 1
	for num >= 1 { // задаем цикл с условием
		num /= 10 // делим данное число на 10 до выполнения условия - num >= 1
		div *= 10 // домножаем каждую итерацию делитель на 10
	}

	return div / 10 // уменьшаем делитель на 10, чтобы иметь возможность вычислить первое(самое левое) значение данного числа
}

func left_digit_finder(num, div int) (int, int) { // находим первое(самое левое) значение данного числа

	left_digit := num / div          // делим данное число на делитель для вычисления первой цифры числа
	rest := num - (left_digit * div) // вычисляем остаток от этого числа
	if num < 0 {                     // проверяем код на наличие ошибки округления целого числа

		rest = num - ((left_digit - 1) * div)

	}
	return left_digit, rest // возвращаем первую цифру данного числа и остаток от этого же числа

}

func SumDigitsRecursive(num int) int { // Вычисляем сумму(sum) всех цифр данного числа
	var sum int

	if num < 0 { // берем модуль числа если оно меншье нуля, чтобы минус не мешал вычислениям

		num *= -1
	}
	if num < 10 { // если число меньше 10 возвращаем его, тк в однозначном числе нечего отделять
		return num
	} else { // вычисляем сумму с помощью рекурсии
		left_digit, rest := left_digit_finder(num, div(num))
		sum = sum + left_digit + SumDigitsRecursive(rest)
		return sum
	}

}

func CalculateDigitalRoot(n int) int {
	if n < 10 { // если число меньше 10 возвращаем его, тк в однозначном числе нечего отделять
		return n
	}

	return CalculateDigitalRoot(SumDigitsRecursive(n))

}

func main() {
	fmt.Println(CalculateDigitalRoot(123456))
}
