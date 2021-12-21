package main

import (
	"bufio"
	"fmt"
	. "gb/lesson10/lesson4/sort"
	"os"
	"strconv"
	"strings"
)

func convertSliceToString(sliceToConvert []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(sliceToConvert), " ", delim, -1), "[]")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter integers to sort splitted by space and press Enter or enter q for exit")
	for scanner.Scan() {
		numbers := []int{}
		input := scanner.Text()

		if strings.TrimSpace(input) == "q" {
			break
		}

		for _, number := range strings.Split(input, " ") {
			number = strings.TrimSpace(number)
			if number != "" {
				var parsedNumber int
				var err error
				if parsedNumber, err = strconv.Atoi(number); err != nil {
					fmt.Println("Error happend during parcing, try again")
					continue
				}
				numbers = append(numbers, parsedNumber)
			}
		}

		Sort(numbers)

		if len(numbers) != 0 {
			fmt.Println("Result:", convertSliceToString(numbers, " "))
		}

		fmt.Println("Enter numbers again or q to exit")
	}
}
