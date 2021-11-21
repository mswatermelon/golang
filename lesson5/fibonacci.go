package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FibCashe map[int64]int64

func getFibonacciNumber(n int64) int64 {
	var fibNumber FibCashe = make(FibCashe)
	return calcFibonacciValue(n, fibNumber)
}

func calcFibonacciValue(n int64, fibCache FibCashe) int64 {
	if _, ok := fibCache[n]; ok {
		return fibCache[n]
	}
	if n <= 1 {
		fibCache[n] = n
		return n
	}

	fibCache[n] = calcFibonacciValue(n-1, fibCache) + calcFibonacciValue(n-2, fibCache)
	return fibCache[n]
}

func main() {
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

			fmt.Println("Your result is:", getFibonacciNumber(parsedNumber))
		}
		fmt.Println("Enter a number again or q to exit")
	}
}
