package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type OperationData struct {
	operator string
	operand1,
	operand2 float64
}

func calculate(operator string, operand1, operand2 float64) (answer string, err error) {
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

func askArguments() (data OperationData, err error) {
	fmt.Println("Please, enter operator")

	if _, err = fmt.Scan(&data.operator); err != nil {
		fmt.Println("Operand is not correct")
		return data, err
	}

	fmt.Println("Please, enter operand 1")

	if _, err = fmt.Scan(&data.operand1); err != nil {
		fmt.Println("First operand is not correct")
		return data, err
	}

	if data.operator != "PRIME" {
		fmt.Println("Please, enter operand 2")

		if _, err = fmt.Scan(&data.operand2); err != nil {
			fmt.Println("Second operand is not correct")
			return data, err
		}
	}

	return
}

func main() {
	var data = OperationData{}
	var err error

	if len(os.Args) >= 2 && strings.Contains(os.Args[1:2][0], "--np") {
		initOperatorFlag := flag.String("np", ",", "Operation to perform")
		initOperand1Flag := flag.Float64("x", 0, "operand 1")
		initOperand2Flag := flag.Float64("y", 0, "operand 2")
		flag.Parse()
		if *initOperatorFlag == "" {
			err = fmt.Errorf("There is nothing to calculate")
		}
		data.operator = *initOperatorFlag
		data.operand1 = *initOperand1Flag
		data.operand2 = *initOperand2Flag
	} else {
		err = fmt.Errorf("There is nothing to calculate")
	}

	if err == nil {
		if result, err := calculate(data.operator, data.operand1, data.operand2); err == nil {
			fmt.Println("Yor result", result)
			return
		}
	}

	for {
		data, err = askArguments()
		if err != nil {
			fmt.Println("Try again")
			fmt.Scanln()
			continue
		}

		result, err := calculate(data.operator, data.operand1, data.operand2)
		if err != nil {
			fmt.Println("Try again")
			fmt.Scanln()
			continue
		}

		fmt.Println("Yor result", result)
		break
	}
}
