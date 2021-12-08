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
