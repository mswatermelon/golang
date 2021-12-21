package main

import (
	"bufio"
	"fmt"
	. "gb/lesson10/lesson5/fib_calculator"
	"os"
	"strconv"
	"strings"
)

func main() {
	calculator := FibCalculator()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a number for which we need to calculate Fibonacci numbers or q to exit")
	for scanner.Scan() {
		input := scanner.Text()
		number := strings.TrimSpace(input)
		if number == "q" {
			break
		}
		if number != "" {
			var parsedNumber int64
			var err error
			if parsedNumber, err = strconv.ParseInt(number, 10, 64); err != nil {
				fmt.Println("Error happend during parcing")
				continue
			}

			fmt.Println("Your result is:", calculator(parsedNumber))
		}
		fmt.Println("Enter a number again or q to exit")
	}
}
