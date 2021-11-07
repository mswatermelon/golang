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
