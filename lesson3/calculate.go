package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

func calculate(operator string, operand1 float64, operand2 float64) (result float64) {
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "":
		result = operand1 * operand2
	case "sqrt":
		result = math.Sqrt(operand1)
	case "/":
		if operand2 == 0 {
			fmt.Println("Forbidden devision on 0")
			os.Exit(1)
		}
		result = operand1 / operand2
	case "PRIME":

	default:
		fmt.Println("Can not calculate this operation")
	}

	return
}

func askArguments() (operator string, operand1, operand2 float64) {
	fmt.Println("Please, enter operand")

	if _, err := fmt.Scan(&operator); err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	fmt.Println("Please, enter operand 1")

	if _, err := fmt.Scan(&operand1); err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	fmt.Println("Please, enter operand 2")

	if _, err := fmt.Scan(&operand2); err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	return operator, operand1, operand2
}

func main() {
	var initOperator string
	var initOperand1, initOperand2 float64

	if len(os.Args) >= 2 && strings.Contains(os.Args[1:2][0], "--np") {
		initOperatorFlag := flag.String("np", ",", "Operation to perform")
		initOperand1Flag := flag.Float64("x", 0, "operand 1")
		initOperand2Flag := flag.Float64("y", 0, "operand 2")
		flag.Parse()
		fmt.Println(os.Args[1:2], initOperator, initOperand1, initOperand2)
		if *initOperatorFlag == "" {
			fmt.Println("Enter an operator")
			os.Exit(1)
		}
		initOperator = *initOperatorFlag
		initOperand1 = *initOperand1Flag
		initOperand2 = *initOperand2Flag
	} else {
		initOperator, initOperand1, initOperand2 = askArguments()
	}

	fmt.Println("Yor result", calculate(initOperator, initOperand1, initOperand2))
}
