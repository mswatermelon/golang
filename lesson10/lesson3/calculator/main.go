package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Calculate(operator string, operand1, operand2 float64) (answer string, err error) {
	var result float64
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "sqrt":
		result = math.Sqrt(operand1)
	case "/":
		if operand2 == 0 {
			fmt.Println("Forbidden devision on 0")
			err = fmt.Errorf("Zero devision")
		}
		result = operand1 / operand2
	case "PRIME":
		var primeNumbersCounter int
		var purposeNumber int = int(math.Floor(operand1))
		var primeResult = make([]int, 0)
		for primeNumbersCounter = 2; primeNumbersCounter <= purposeNumber; primeNumbersCounter++ {
			var flag bool = false
			for j := 2; j <= primeNumbersCounter/2; j++ {
				if primeNumbersCounter%j == 0 {
					flag = true
				}
			}
			if flag == false {
				primeResult = append(primeResult, primeNumbersCounter)
			}
		}
		answer = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(primeResult)), ","), "[]")
		return
	default:
		fmt.Println("Can not calculate this operation")
		err = fmt.Errorf("Wrong operator")
	}

	intPart, fractPart := math.Modf(result)

	if fractPart == 0 {
		answer = strconv.FormatInt(int64(intPart), 10)
		return
	}

	answer = strconv.FormatFloat(result, 'f', 6, 64)
	return
}
