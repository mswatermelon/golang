package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Hello, let us analize a three-digit number")

	var number float64 // or var number uint64

	fmt.Println("Please, enter the three-digit number")

	if _, err := fmt.Scan(&number); err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	if number < 100 || number > 999 {
		fmt.Println("Error")
		os.Exit(1)
	}

	n1 := math.Floor(number / 100)
	n2 := math.Floor(((number - n1*100) / 10))
	n3 := math.Floor(number - n1*100 - n2*10)

	// or if uint64
	// n1 := number / 100
	// n2 := number / 100 % 10
	// n3 := number % 10

	fmt.Println("Number of hundreds", n1, "number of tens", n2, "number of units", n3)
}
