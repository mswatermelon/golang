// 1. Проанализируйте задания предыдущих уроков.
// a. В каких случаях необходима была явная передача указателя в качестве входных параметров и возвращаемых результатов или в качестве приёмника в методах? 
// - конечно же в fmt.Scan, fmt.Scanln
// - расписала в файле где можно было бы использовать ещё
// b. В каких случаях мы фактически имеем дело с указателями при передаче параметров, хотя явно их не указываем?
// - в случае со слайсами
// - в случае с мапами
// - в случае с defer, переменные при инициализации захватываются по указателю

// ДЗ1 - не нашла мест, где можно было бы заменить на указатель
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
}

// ДЗ 2 - не нашла мест, где можно было бы заменить на указатель
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, this is a program for calculating an area of a rectangle")

	var operand1, operand2 float64

	fmt.Println("Please, enter the length of the first side of the rectangle")

	if _, err := fmt.Scan(&operand1); err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	fmt.Println("Please, enter the length of the second side of the rectangle")

	if _, err := fmt.Scan(&operand2); err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	fmt.Println("This is your result", operand1*operand2)
}

// ДЗ 3
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

// Тут могла бы заменить на структуру и передавать ее по ссылке
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

// Возвращаемое значение могло быть ссылкой, так как это структура
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

// ДЗ 4 - не нашла мест, где можно было бы заменить на указатель
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertSliceToString(sliceToConvert []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(sliceToConvert), " ", delim, -1), "[]")
}

func sort(numbers []int64) {
	for i := 1; i < len(numbers); i++ {
		number := numbers[i]
		j := i
		for ; j >= 1 && numbers[j-1] > number; j-- {
			numbers[j] = numbers[j-1]
		}
		numbers[j] = number
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter integers to sort splitted by space and press Enter or enter q for exit")
	for scanner.Scan() {
		numbers := []int64{}
		input := scanner.Text()

		if strings.TrimSpace(input) == "q" {
			break
		}

		for _, number := range strings.Split(input, " ") {
			number = strings.TrimSpace(number)
			if number != "" {
				var parsedNumber int64
				var err error
				if parsedNumber, err = strconv.ParseInt(number, 10, 64); err != nil {
					fmt.Println("Error happend during parcing, try again")
					continue
				}
				numbers = append(numbers, parsedNumber)
			}
		}

		sort(numbers)

		if len(numbers) != 0 {
			fmt.Println("Result:", convertSliceToString(numbers, " "))
		}

		fmt.Println("Enter numbers again or q to exit")
	}
}
// 2. Для арифметического умножения и разыменования указателей в Go используется один и тот же символ — оператор (*).
// Как вы думаете, как компилятор Go понимает, в каких случаях в выражении имеется в виду умножение, а в каких — разыменование указателя?
// Поправила старый ответ:
// Думаю, что он смотрит что находится рядом, если справа строка, то он ищет ее среди ключевых слов, а затем среди наименований переменных.
// Если слева и справа числа, то производит математическую операцию
// Новый:
// 1.Компилятор видит * - он знает, что это либо умножение, либо разыменование. Какое принципиальное отличие,
// по которому компилятор сразу поймет что подразумевалось - умножение или разыменование
// - Смотрит что слева, если переменная или число, строка (операнды), то сразу умножение
// 2. Компилятор видит конструкцию fmt.Println(a *** b) - что ему нужно проверить , чтобы понять - валидна ли она, выполнима ли
// a) Проверить тип a. Eсли это int, float, complex, то продолжить. Если нет - ошибка.
// b) Проверить тип **b. Если тип не совпадает с типом a - ошибка.
