package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

type Figure interface {
	area() float64
}

type Rectangle struct {
	a float64
	b float64
}

func (rec Rectangle) area() float64 {
	return rec.a*rec.b
}

type Square struct {
	a float64
}

func (square Square) area() float64 {
	return square.a*square.a
}

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

var ScanError = errors.New("Scanning error")

func scanOperand(operandNumber string, operand *float64) error {
	fmt.Println("Please, enter the length of the", operandNumber, "side")
	if _, err := fmt.Scan(operand); err != nil {
		return fmt.Errorf("%w: error on scanning first operand", ScanError)
	}

	return nil
}

func scanFigure(figure *string) error {
	fmt.Println("Please, choose a figure type, which you want to calculate area for: r - rectangle, s - square, c - circle")
	if _, err := fmt.Scan(figure); err != nil {
		return fmt.Errorf("%w: error on scanning first operand", ScanError)
	}

	return nil
}

func decideWhichType(figureType string, a, b float64) (Figure, error) {
	switch figureType {
	case "c":
		return Circle{radius: a}, nil
	case "r":
		return Rectangle{a: a, b: b}, nil
	case "s":
		return Square{a: a}, nil
	}
	return nil, fmt.Errorf("You made a mistake when selected a figure type")
}

func calcAreaDependingOnType(figureType string, a float64, b float64) (float64, error) {
	figure, err := decideWhichType(figureType, a, b)
	if err != nil {
		return 0, err
	}
	return figure.area(), nil
}

func main() {
	fmt.Println("Hello, this is a program for calculating an area")

	var figureType string
	if err := scanFigure(&figureType); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Enter figure's side length or radius for circle")
	var operand1, operand2 float64
	if err := scanOperand("first", &operand1); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if figureType == "r" {
		if err := scanOperand("second", &operand2); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var result float64
	var err error

	if result, err = calcAreaDependingOnType(figureType, operand1, operand2); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("This is your result", result)
}
