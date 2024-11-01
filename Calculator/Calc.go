package main

import (
	"fmt"
	"strconv"
	"strings"
)

const IsLeftParenthesis = 1
const IsRightParenthesis = 2
const IsNotParenthesis = 0
const IsMultiplication = 10
const IsDivision = 20
const IsAddition = 30
const IsSubtraction = 40
const IsNotOperation = 0
const IsPoint = 100
const IsNotSeparator = 0

func IsNumber(char byte) bool {
	const numbers = "1234567890"
	for index, _ := range numbers {
		if numbers[index] == char {
			return true
		}
	}
	return false
}

func IsParenthesis(char byte) int {
	if string(char) == "(" {
		return IsLeftParenthesis
	}
	if string(char) == ")" {
		return IsRightParenthesis
	}
	return IsNotParenthesis
}

func IsOperator(char byte) int {
	if string(char) == "*" {
		return IsMultiplication
	} else if string(char) == "/" {
		return IsDivision
	} else if string(char) == "+" {
		return IsAddition
	} else if string(char) == "-" {
		return IsSubtraction
	}
	return IsNotOperation
}

func IsSeparator(char byte) int {
	if string(char) == "." {
		return IsPoint
	}
	return IsNotSeparator
}

func GetPryority(operator int) int {
	mapofoperators := map[int]int{
		IsMultiplication: 2,
		IsDivision:       2,
		IsAddition:       1,
		IsSubtraction:    1,
	}
	pryority := mapofoperators[operator]
	return pryority
}

func ExtractNum(Expression string, indexofnum int) (string, int) {
	var num string
	var index int
	var length int = len(Expression)
	for nextnotnumindex := indexofnum; nextnotnumindex < length; nextnotnumindex++ {
		if IsNumber(Expression[nextnotnumindex]) || IsSeparator(Expression[nextnotnumindex]) != 0 {
			num += string(Expression[nextnotnumindex])
		}
		if !IsNumber(Expression[nextnotnumindex]) && IsSeparator(Expression[nextnotnumindex]) == 0 {
			return num, nextnotnumindex
		}
		index = nextnotnumindex
	}
	return num, index
}

func PopNum(sliceofnums []float64, numtopop int) ([]float64, []float64, error) {

	var poppednum, newsliceofnums []float64

	if numtopop > len(sliceofnums) {
		return poppednum, sliceofnums, fmt.Errorf("numtopop > length of slice of nums, %d", numtopop)
	}
	if numtopop <= 0 {
		return poppednum, sliceofnums, fmt.Errorf("numtopop <= 0, %d", numtopop)
	}

	poppednum = append(sliceofnums[len(sliceofnums)-numtopop:])
	newsliceofnums = append(sliceofnums[:len(sliceofnums)-numtopop], sliceofnums[len(sliceofnums):]...)

	return poppednum, newsliceofnums, nil
}

func PopOp(opslice []int) (int, []int, error) {

	var newopslice []int

	if len(opslice) == 0 {
		return 0, opslice, fmt.Errorf("no operator to pop")
	}

	poppedop := opslice[len(opslice)-1]
	newopslice = append(opslice[:len(opslice)-1], opslice[len(opslice):]...)

	return poppedop, newopslice, nil
}

func Transact(sliceofnums []float64, opslice []int, operator int, addop bool) (float64, []float64, []int, error) {
	var result float64
	var poppedop int
	var poppednums []float64
	var popnumerr, popoperr error

	poppedop, opslice, popoperr = PopOp(opslice)
	poppednums, sliceofnums, popnumerr = PopNum(sliceofnums, 2)

	if poppedop == 0 && popoperr != nil {
		return 0, sliceofnums, opslice, popoperr
	}

	if poppednums != nil && popnumerr != nil {
		return 0, sliceofnums, opslice, popnumerr
	}

	switch {
	case poppedop == IsAddition:
		result = poppednums[0] + poppednums[1]
	case poppedop == IsSubtraction:
		result = poppednums[0] - poppednums[1]
	case poppedop == IsMultiplication:
		result = poppednums[0] * poppednums[1]
	case poppedop == IsDivision:
		if poppednums[1] == 0 {
			return 0, sliceofnums, opslice, fmt.Errorf("division by zero")
		}
		result = poppednums[0] / poppednums[1]
	}

	sliceofnums = append(sliceofnums, result)
	if addop {
		opslice = append(opslice, operator)
	}

	return result, sliceofnums, opslice, nil
}

func IsCorrectExpression(Expression string) (bool, error) { //Проверка на правильность заданной строки
	Expression = strings.Replace(Expression, " ", "", -1)
	if Expression == "" {
		return false, fmt.Errorf("empty expression")
	}
	correctexpression := true
	expressionlength := len(Expression)
	countleftparenthesis := 0
	countrightparenthesis := 0
	for index, _ := range Expression {
		if index < expressionlength-1 {
			switch {
			case !IsNumber(Expression[index]) && IsParenthesis(Expression[index]) == 0 && IsOperator(Expression[index]) == 0 && IsSeparator(Expression[index]) == 0: //Недопустимые символы
				correctexpression = false
				//fmt.Printf("incorrect symbol, char %d. Allowed only: %s ", index, "1234567890.*/+-()")
			case index == 0 && !IsNumber(Expression[index]) && IsParenthesis(Expression[index]) == 0: //Запрещенная последовательность "выражение начинается не числом и не скобкой"
				correctexpression = false
				//fmt.Printf(`wrong sequence "non-number character": char %d `, index)
			case IsOperator(Expression[index]) != 0 && IsOperator(Expression[index+1]) != 0: //Запрещенная последовательность "оператор->оператор"
				correctexpression = false
				//fmt.Printf(`wrong sequence "operation sign->operation sign": chars %d, %d `, index, index+1)
			case IsSeparator(Expression[index]) != 0 && IsSeparator(Expression[index+1]) != 0: //Запрещенная последовательность "разделитель->разделитель"
				correctexpression = false
				//fmt.Printf(`wrong sequence "multiple separators together": starting from char %d `, index)
			case IsParenthesis(Expression[index]) != 0 && IsSeparator(Expression[index+1]) != 0: //Запрещенная последовательность "скобка->разделитель дроби"
				correctexpression = false
				//fmt.Printf(`wrong sequence "parenthesis->separator": chars %d, %d `, index, index+1)
			case IsParenthesis(Expression[index+1]) != 0 && IsSeparator(Expression[index]) != 0: //Запрещенная последовательность "разделитель дроби->скобка"
				correctexpression = false
				//fmt.Printf(`wrong sequence "separator->parenthesis": chars %d, %d `, index, index+1)
			case IsSeparator(Expression[index]) != 0 && IsOperator(Expression[index+1]) != 0: //Запрещенная последовательность "разделитель дроби->оператор
				correctexpression = false
				//fmt.Printf(`wrong sequence "separator->operation sign": chars %d, %d `, index, index+1)
			case IsSeparator(Expression[index+1]) != 0 && IsOperator(Expression[index]) != 0: //Запрещенная последовательность "оператор->разделитель дроби"
				correctexpression = false
				//fmt.Printf(`wrong sequence "operation sign->separator": chars %d, %d `, index, index+1)
			case IsSeparator(Expression[index]) != 0 && IsNumber(Expression[index+1]) && IsNumber(Expression[index-1]): //Запрещенная последовательность "множественные разделители дроби в числе"
				for nextcharindex := index + 1; nextcharindex < expressionlength; nextcharindex++ {
					if !IsNumber(Expression[nextcharindex]) {
						if IsSeparator(Expression[nextcharindex]) != 0 {
							correctexpression = false
							//fmt.Printf(`wrong sequence "multiple separators within number": starting from char %d `, index)
							break
						} else {
							break
						}
					}
				}
			case IsParenthesis(Expression[index]) == IsLeftParenthesis && IsParenthesis(Expression[index+1]) == IsRightParenthesis: //Запрещенная последовательность "пустые скобки"
				correctexpression = false
				//fmt.Printf(`wrong sequence "empty parentheses": chars %d, %d `, index, index+1)
			case IsParenthesis(Expression[index]) == IsRightParenthesis && countleftparenthesis == 0: // Запрещенная последовательность "подвыражение начинается с правой скобки"
				countrightparenthesis++
				correctexpression = false
				//fmt.Printf(`wrong sequence "beginning form right parenthesis": on char %d `, index)
			case IsParenthesis(Expression[index]) == IsLeftParenthesis && countleftparenthesis == 0: // Считаем левые и правые скобки
				countleftparenthesis++
				for nextcharindex := index + 1; nextcharindex < expressionlength; nextcharindex++ {
					if IsParenthesis(Expression[nextcharindex]) == IsLeftParenthesis {
						countleftparenthesis++
					} else if IsParenthesis(Expression[nextcharindex]) == IsRightParenthesis {
						countrightparenthesis++
					}

				}
			}
		} else if !IsNumber(Expression[index]) && IsParenthesis(Expression[index]) == 0 && IsOperator(Expression[index]) == 0 && IsSeparator(Expression[index]) == 0 { //Недопустимые символы
			correctexpression = false
			//fmt.Printf("incorrect symbol, char %d. Allowed only: %s", index, "1234567890.*/+-()")
		} else if !IsNumber(Expression[index]) && index == expressionlength-1 {
			correctexpression = false
			//fmt.Printf(`wrong sequence "non-numeric last character"`)
		}
	}

	if countleftparenthesis < countrightparenthesis { // Не хватает левых скобок
		correctexpression = false
		//fmt.Printf(`wrong sequence "insufficient number of left parentheses"`)
	} else if countleftparenthesis > countrightparenthesis { // Не хватает правых скобок
		correctexpression = false
		//fmt.Printf(`wrong sequence "insufficient number of right parentheses"`)
	}

	if !correctexpression {
		return false, fmt.Errorf("incorrect expression")
	}
	return true, nil
}

func TokenizeandCalc(Expression string) (float64, error) {
	var result float64
	var operatorsslice []int
	var numsslice []float64
	var priority, countdown int
	var matherr error

	length := len(Expression)
	for indexoftokenizer := 0; indexoftokenizer < length; indexoftokenizer++ {
		operatorslicelength := len(operatorsslice)
		if IsNumber(Expression[indexoftokenizer]) {
			num, afternumberindex := ExtractNum(Expression, indexoftokenizer)
			indexoftokenizer = afternumberindex
			numinfloat, _ := strconv.ParseFloat(num, 8)
			numsslice = append(numsslice, numinfloat)
		}
		if !IsNumber(Expression[indexoftokenizer]) && IsSeparator(Expression[indexoftokenizer]) == 0 {
			switch {
			case IsOperator(Expression[indexoftokenizer]) != 0:
				if operatorslicelength-1 >= 0 {
					priority = GetPryority(IsOperator(Expression[indexoftokenizer]))
					if GetPryority(operatorsslice[operatorslicelength-1]) == priority {
						result, numsslice, operatorsslice, matherr = Transact(numsslice, operatorsslice, IsOperator(Expression[indexoftokenizer]), true)
						if result == 0 && matherr != nil {
							return 0, matherr

						}
					} else if GetPryority(operatorsslice[operatorslicelength-1]) < priority {
						operatorsslice = append(operatorsslice, IsOperator(Expression[indexoftokenizer]))
					} else if GetPryority(operatorsslice[operatorslicelength-1]) > priority {
						result, numsslice, operatorsslice, matherr = Transact(numsslice, operatorsslice, IsOperator(Expression[indexoftokenizer]), true)
						if result == 0 && matherr != nil {
							return 0, matherr

						}
					}

				} else {
					operatorsslice = append(operatorsslice, IsOperator(Expression[indexoftokenizer]))
				}
			case IsParenthesis(Expression[indexoftokenizer]) == IsLeftParenthesis:
				operatorsslice = append(operatorsslice, IsLeftParenthesis)
			case IsParenthesis(Expression[indexoftokenizer]) == IsRightParenthesis:
				for {
					if (operatorsslice[len(operatorsslice)-1]) == IsLeftParenthesis {
						_, operatorsslice, _ = PopOp(operatorsslice)
						break
					}
					result, numsslice, operatorsslice, matherr = Transact(numsslice, operatorsslice, 0, false)
					if result == 0 && matherr != nil {
						return 0, matherr

					}
				}
			}
		}
	}

	countdown = len(operatorsslice) - 1
	for {
		if countdown < 0 {
			break
		} else {
			result, numsslice, operatorsslice, matherr = Transact(numsslice, operatorsslice, 0, false)
			if result == 0 && matherr != nil {
				return 0, matherr

			}
		}
		countdown--
	}
	return numsslice[0], nil
}

func Calc(Expression string) (float64, error) {

	check, checkerr := IsCorrectExpression(Expression)
	if check && checkerr == nil {
		result, calcerr := TokenizeandCalc(Expression)
		if result == 0 && calcerr != nil {
			return 0, calcerr
		} else {
			return result, nil
		}

	} else {
		return 0, checkerr
	}
}

func main() {
	fmt.Println(Calc("1*2.54+41+((3/3+10)/2-2.5*10)-1"))
	fmt.Println(Calc("1+1"))
	fmt.Println(Calc("(2+2)*2"))
	fmt.Println(Calc("2+2*2")) //? 6
	fmt.Println(Calc("1/2"))
	fmt.Println(Calc("1+1*"))
	fmt.Print(Calc(""))
}
