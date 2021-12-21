package main_test

import (
	. "gb/lesson10/lesson5/fib_calculator"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcFibonacciValueZero(t *testing.T) {
	var fibNumber FibCache = make(FibCache)
	if result := CalcFibonacciValue(0, fibNumber); result != 0 {
		t.Errorf("Fibonacci(0) expected to be 0, but received %d", result)
	}
}

func TestCalcFibonacciValueOne(t *testing.T) {
	var fibNumber FibCache = make(FibCache)
	if result := CalcFibonacciValue(1, fibNumber); result != 1 {
		t.Errorf("Fibonacci(0) expected to be 0, but received %d", result)
	}
}

func TestCalcFibonacciValueNotToBeZero(t *testing.T) {
	var fibNumber FibCache = make(FibCache)

	for i := 1; i < 100; i++ {
		randI := rand.Int63n(40) + 1
		if result := CalcFibonacciValue(randI, fibNumber); result == 0 {
			t.Errorf("Fibonacci(0) expected NOT to be 0, but received %d", result)
		}
	}
}

func TestCalcFibonacciValueFull(t *testing.T) {
	var fibNumber FibCache = make(FibCache)
	fibonacciNumbers := []int64{
		0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811,
	}

	for i, expected := range fibonacciNumbers {
		got := CalcFibonacciValue(int64(i), fibNumber)

		if expected != got {
			t.Errorf("Fibonacci(%d) expected to be %d, but received %d", i, expected, got)
		}
	}
}

func BenchmarkCalcFibonacciValue(b *testing.B) {
	var fibNumber FibCache = make(FibCache)
	for i := 0; i < b.N; i++ {
		CalcFibonacciValue(int64(i), fibNumber)
	}
}

// testify
func TestCalcFibonacciValueWithPack(t *testing.T) {
	var fibNumber FibCache = make(FibCache)
	result := CalcFibonacciValue(0, fibNumber)
	// assert equality
	assert.Equal(t, result, int64(0), "Fibonacci(0) expected to be 0, but received %d", result)
}

// Удобно, что предусмотрен вывод ошибок и даже с аргументами
// Получается два действия - сравнение и вывод ошибки объеденены в одно
// Удобно, что код в месте проверки будет единообразен, намного легче найти место,
// где это происходит среди вычислений
