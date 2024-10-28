
func extractnumbers(Expression string) string {

	for index, _ := range Expression {
		if !IsNumber(Expression[index]) {
			if IsSeparator(Expression[index]) == 0 {
				if IsParenthesis(Expression[index]) != 0 {
					continue
				}
				outputstring += ","
			}
		}
		if IsSeparator(Expression[index]) != 0 {
			outputstring += "."
		}
		if IsNumber(Expression[index]) {
			outputstring += string(Expression[index])
		}
	}
	return outputstring
}

func Tokenizer(Expression string) (string, error) { // Парсинг выражения

	for index, _ := range Expression {
		if !IsNumber(Expression[index]) && IsSeparator(Expression[index]) == 0 {
			for indexofstack, _ := range thesliceforstack {
				switch {
				case string(Expression[index]) == "+":
					if thesliceforstack[0] == "-" {
						outputstring = outputstring + "," + thesliceforstack[indexofstack]
						thesliceforstack[indexofstack] = ""
					}
				case string(Expression[index]) == "-":
					if thesliceforstack[0] == "+" {
						outputstring = outputstring + "," + thesliceforstack[indexofstack]
						thesliceforstack[indexofstack] = ""
					}
				case string(Expression[index]) == "*":
					if thesliceforstack[0] == "//" || thesliceforstack[0] == "-" || thesliceforstack[0] == "+" {
						outputstring = outputstring + "," + thesliceforstack[indexofstack]
						thesliceforstack[0] = ""
					}
				case string(Expression[index]) == "//":
					if thesliceforstack[0] == "*" || thesliceforstack[0] == "-" || thesliceforstack[0] == "+" {
						outputstring = outputstring + "," + thesliceforstack[indexofstack]
						thesliceforstack[indexofstack] = ""
					}
				default:
					continue
				}
			}
			thesliceforstack = append(thesliceforstack, string(Expression[index]))
		} else {
			outputstring += "1"
		}

	}

	stack := strings.Join(thesliceforstack, "")
	outputstring += stack
	fmt.Println(outputstring)
	return outputstring, nil
}

func Calculator(Expression string) (float64, error) { //Выполнение элементарных операций

	// (devide / 0)

	for indexofcalculator, _ := range Expression {
		if !IsNumber(Expression[indexofcalculator]) {
			//sum := Expression[indexofcalculator*-1] + Expression[indexofcalculator] + Expression[indexofcalculator*-1-1]
			//sum = int(sum)
		}
	}
	return 1, nil
}

func main() {
	//fmt.Println(IsCorrectExpression("1*2.5441+(3/3)-2.5("))
	Tokenizer("1*2.5441+(3/3)-2.5")
}

/*
	else {
			for {
				if !IsNumber(Expression[index]) {
					if IsSeparator(Expression[index]) == 0 {
						break
					} else if index > 0 {
						outputstring = outputstring + string(Expression[index-1])
						break
					} else if index == 0 {
						outputstring = outputstring + string(Expression[index])
					}
				}
				outputstring = outputstring + string(Expression[index])
				if index == expressionlength-1 {
					break
				}
				index++
			}
			outputstring += ","
		}
	}
*/

/*
func extractthefirstnumber(Expression string) string {
	var firstnumber string
	var reexp string

	for indexofthefirstnumber, _ := range Expression {
		if IsNumber(Expression[indexofthefirstnumber]) {
			firstnumber += string(Expression[indexofthefirstnumber])
		}
		if !IsNumber(Expression[indexofthefirstnumber]) {
			if IsSeparator(Expression[indexofthefirstnumber]) != 0 {
				firstnumber += "."
			} else {
				reexp = strings.Replace(Expression, firstnumber, ",", -1)
				return reexp
			}

		}
	}
	return ""
}

func reshapeexpression(reex string) string {

	reexp = strings.Replace(reexp, num1, "", -1)

	if !IsNumber(reexp[0]) {
		reexp = strings.Replace(reexp, string(reexp[0]), "", -1)
		if !IsNumber(reexp[1]) {
			reexp = strings.Replace(reexp, string(reexp[1]), "", -1)
		}
	return reex
}
}

func extractnums(reexp, num2 string) string {
	var num1 string


		for indexofnum1, _ := range reexp {
			if IsNumber(reexp[indexofnum1]) {

				num1 += string(reexp[indexofnum1])
			}
			if !IsNumber(reexp[indexofnum1]) {
				if IsSeparator(reexp[indexofnum1]) != 0 {
					num1 += "."
				} else {
					return num1
				}
			}
		}
		return ""
	}


func extractnum2(reexp, num1 string) string {
	var num2 string

	reexp = strings.Replace(reexp, num1, "", -1)

	if !IsNumber(reexp[0]) {
		reexp = strings.Replace(reexp, string(reexp[0]), "", -1)
		if !IsNumber(reexp[1]) {
			reexp = strings.Replace(reexp, string(reexp[1]), "", -1)
		}

		for indexofnum2, _ := range reexp {
			if IsNumber(reexp[indexofnum2]) {

				num2 += string(reexp[indexofnum2])
			}
			if !IsNumber(reexp[indexofnum2]) {
				if IsSeparator(reexp[indexofnum2]) != 0 {
					num2 += "."
				} else {
					return num2
				}
			}
		}
	}
	return ""
}
*/

/*
	else {
			for {
				if !IsNumber(Expression[index]) {
					if IsSeparator(Expression[index]) == 0 {
						break
					} else if index > 0 {
						outputstring = outputstring + string(Expression[index-1])
						break
					} else if index == 0 {
						outputstring = outputstring + string(Expression[index])
					}
				}
				outputstring = outputstring + string(Expression[index])
				if index == expressionlength-1 {
					break
				}
				index++
			}
			outputstring += ","
		}
	}
*/

/*
func extractthefirstnumber(Expression string) string {
	var firstnumber string
	var reexp string

	for indexofthefirstnumber, _ := range Expression {
		if IsNumber(Expression[indexofthefirstnumber]) {
			firstnumber += string(Expression[indexofthefirstnumber])
		}
		if !IsNumber(Expression[indexofthefirstnumber]) {
			if IsSeparator(Expression[indexofthefirstnumber]) != 0 {
				firstnumber += "."
			} else {
				reexp = strings.Replace(Expression, firstnumber, ",", -1)
				return reexp
			}

		}
	}
	return ""
}

func reshapeexpression(reex string) string {

	reexp = strings.Replace(reexp, num1, "", -1)

	if !IsNumber(reexp[0]) {
		reexp = strings.Replace(reexp, string(reexp[0]), "", -1)
		if !IsNumber(reexp[1]) {
			reexp = strings.Replace(reexp, string(reexp[1]), "", -1)
		}
	return reex
}
}

func extractnums(reexp, num2 string) string {
	var num1 string


		for indexofnum1, _ := range reexp {
			if IsNumber(reexp[indexofnum1]) {

				num1 += string(reexp[indexofnum1])
			}
			if !IsNumber(reexp[indexofnum1]) {
				if IsSeparator(reexp[indexofnum1]) != 0 {
					num1 += "."
				} else {
					return num1
				}
			}
		}
		return ""
	}

func extractnum2(reexp, num1 string) string {
	var num2 string

	reexp = strings.Replace(reexp, num1, "", -1)

	if !IsNumber(reexp[0]) {
		reexp = strings.Replace(reexp, string(reexp[0]), "", -1)
		if !IsNumber(reexp[1]) {
			reexp = strings.Replace(reexp, string(reexp[1]), "", -1)
		}

		for indexofnum2, _ := range reexp {
			if IsNumber(reexp[indexofnum2]) {

				num2 += string(reexp[indexofnum2])
			}
			if !IsNumber(reexp[indexofnum2]) {
				if IsSeparator(reexp[indexofnum2]) != 0 {
					num2 += "."
				} else {
					return num2
				}
			}
		}
	}
	return ""
}

for index, _ := range Expression {
		if !IsNumber(Expression[index]) {
			if IsSeparator(Expression[index]) == 0 {
				if IsParenthesis(Expression[index]) != 0 {
					thesliceforstack = append(thesliceforstack, string(Expression[index]))
					continue
				}
				if IsOperator(Expression[index]) != 0 {
					outputstring += ","
					for index1, _ := range thesliceforstack {
						var length int = len(thesliceforstack)
						switch {
						case string(Expression[index]) == "+" && string(thesliceforstack[length-1]) == "-":
							//outputstring += "," + string(thesliceforstack[index])
							thesliceforstack[index1] = ""

						case string(Expression[index]) == "-" && string(thesliceforstack[length-1]) == "+":
							outputstring += "," + string(thesliceforstack[index])
							thesliceforstack[index1] = ""

						case string(Expression[index]) == "*" && (string(thesliceforstack[length-1]) == "//" || string(thesliceforstack[length-1]) == "-" || string(thesliceforstack[length-1]) == "+"):
							outputstring += "," + string(thesliceforstack[index])
							thesliceforstack[index1] = ""

						case string(Expression[index]) == "//" && (string(thesliceforstack[length-1]) == "*" || string(thesliceforstack[length-1]) == "-" || string(thesliceforstack[length-1]) == "+"):
							outputstring += "," + string(thesliceforstack[index])
							thesliceforstack[index1] = ""
						default:
							continue
						}
					}
					thesliceforstack = append(thesliceforstack, string(Expression[index]))
				}
			}
		}

		if IsSeparator(Expression[index]) != 0 {
			outputstring += "."
		}
		if IsNumber(Expression[index]) {
			outputstring += string(Expression[index])
		}
	}
	return outputstring, thesliceforstack
}

func Tokenizer(Expression string) (string, error) { // Парсинг выражения

	for index, _ := range Expression {
		if !IsNumber(Expression[index]) && IsSeparator(Expression[index]) == 0 {
		}
	}

	extractnumbers(Expression)

	stack := strings.Join(thesliceforstack, "")
	outputstring += stack
	fmt.Println(outputstring)
	return outputstring, nil
}
*/
