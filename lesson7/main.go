//1. С какими интерфейсами мы уже сталкивались в предыдущих уроках? Обратите внимание на уроки, в которых мы читали из стандартного ввода и писали в стандартный вывод.
// - пустой интерфейс interface{} (fmt.PrintLn, fmt.Scan, fmt.Errorf)
// - интерфейс error (errors.New)
// - io.Reader, io.Writer
// - FlagSet из библиотеки "flag"

//2. Посмотрите примеры кода в своём портфолио. Везде ли ошибки обрабатываются грамотно? Хотите ли вы переписать какие-либо функции?
// Да, таких мест полно :)
// ДЗ 2
// 1. Я бы разделила сканирование данных и обработку ошибок оставила бы в main.
// С интерфейсами сделала в advanced.go
// Исходный вариант - https://github.com/mswatermelon/golang/pull/2/files#diff-9386e6e0542f9721d315ac58529a209ee154a3adf7839432d6737b0231af53bb
// Новый:
package main

import (
	"fmt"
	"os"
	"errors"
)

var ScanError = errors.New("Scanning error")

func scanOperand(operandNumber string, operand *float64) (err error) {
	fmt.Println("Please, enter the length of the", operandNumber, "side of the rectangle")
	if _, err := fmt.Scan(operand); err != nil {
		return fmt.Errorf("%w: error on scanning first operand", ScanError)
	}

	return nil
}

func main() {
	fmt.Println("Hello, this is a program for calculating an area of a rectangle")

	var operand1, operand2 float64
	if err := scanOperand("first", &operand1); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := scanOperand("second", &operand2); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("This is your result", operand1*operand2)
}
// 2. То же самое, разделилила бы логику чтения площади и обработки ошибки
// + перенесла бы вычисления в функции
// Исходный вариант - https://github.com/mswatermelon/golang/pull/2/files#diff-07087e02ed0d64435d4b29ee2e075e0625465763fe64a44b9352971c4d4c5498
// Новый:
package main

import (
	"fmt"
	"math"
	"os"
)

func scanArea(area *float64) (err error) {
	fmt.Println("Please, enter the area of the circumference")
	if _, err := fmt.Scan(area); err != nil {
		return fmt.Errorf("Number is not correct")
	}

	return nil
}

func getDiameter(area float64) float64 {
	return math.Sqrt(4*area/math.Pi)
}

func getLength(area float64) float64 {
	return math.Sqrt(area*4*math.Pi)
}

func main() {
	fmt.Println("Hello, here we want to calculate circumference's diameter and length")

	var area float64

	if err := scanArea(&area); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("So, the diameter is", getDiameter(area), "and the length is", getLength(area))
}

// 3. Опять разделилила бы логику чтения числа и обработки ошибки, вынесу проверку на трехзначность в функцию
// + поправлю комментарии, что не сделала сразу
// Исходный вариант - https://github.com/mswatermelon/golang/pull/2/files#diff-e09cad00a217ddff3afd96693091ebf66618388170b7fd04b31fdedb7120088c
// Новый
package main

import (
	"fmt"
	"math"
	"os"
)

func scanNumber(number *uint64) (err error) {
	fmt.Println("Please, enter the three-digit number")
	if _, err := fmt.Scan(number); err != nil {
		return fmt.Errorf("Number is not correct")
	}

	return nil
}

func isThreeDigitNumber(number uint64) error {
	if number < 100 || number > 999 {
		return fmt.Errorf("Number is not three-digit")
	}
	return nil
}


func main() {
	fmt.Println("Hello, let us analize a three-digit number")

	var number uint64

	if err := scanNumber(&number); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := isThreeDigitNumber(number); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ones := number / 100
	tens := number / 100 % 10
	hundreds := number % 10

	fmt.Println("Number of hundreds", ones, "number of tens", tens, "number of units", hundreds)
}
