package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Hello, here we want to calculate circumference's diameter and length")

	var area float64

	fmt.Println("Please, enter the area of the circumference")

	if _, err := fmt.Scan(&area); err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	fmt.Println("So, the diameter is", math.Sqrt(4*area/math.Pi), "and the length is", math.Sqrt(area*4*math.Pi))
}
