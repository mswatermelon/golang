package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter integers to sort splitted by space and press Enter")
	for scanner.Scan() {
		numbers := []int64{}
		input := scanner.Text()

		for _, number := range strings.Split(input, " ") {
			number = strings.TrimSpace(number)
			if number != "" {
				var parsedNumber int64
				var err error
				if parsedNumber, err = strconv.ParseInt(number, 10, 64); err != nil {
					fmt.Println("Error happend during parcing")
					os.Exit(1)
				}
				numbers = append(numbers, parsedNumber)
			}
		}

		for i := 1; i < len(numbers); i++ {
			number := numbers[i]
			numberToCompare := i
			for ; numberToCompare >= 1 && numbers[numberToCompare-1] > number; numberToCompare-- {
				numbers[numberToCompare] = numbers[numberToCompare-1]
			}
			numbers[numberToCompare] = number
		}

		if len(numbers) != 0 {
			fmt.Println("Result:")
		}
		for i := 0; i < len(numbers); i++ {
			if i == (len(numbers) - 1) {
				fmt.Print(numbers[i])
				break
			}
			fmt.Print(numbers[i], " ")
		}
		if len(numbers) != 0 {
			break
		}
	}
}
