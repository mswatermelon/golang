package main_test

import (
	"fmt"

	. "gb/lesson10/lesson3/calculator"
)

func ExampleCalculateSum() {
	result, _ := Calculate("+", 7, 9)
	fmt.Println(result)
	// Output: 16
}

func ExampleCalculateMinus() {
	result, _ := Calculate("-", 7, 9)
	fmt.Println(result)
	// Output: -2
}

func ExampleCalculateDevision() {
	result, _ := Calculate("/", 45, 5)
	fmt.Println(result)
	// Output: 9
}

func ExampleCalculateMultiplication() {
	result, _ := Calculate("*", 7, 9)
	fmt.Println(result)
	// Output: 63
}

func ExampleCalculateSqrt() {
	result, _ := Calculate("sqrt", 16, 0)
	fmt.Println(result)
	// Output: 4
}

func ExampleCalculatePrimeNumbers() {
	result, _ := Calculate("PRIME", 16, 0)
	fmt.Println(result)
	// Output: 2,3,5,7,11,13
}
