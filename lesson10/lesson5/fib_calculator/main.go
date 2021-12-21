package fib_calculator

type FibCache map[int64]int64

func GetFibonacciNumber(n int64) int64 {
	var fibNumber FibCache = make(FibCache)
	return CalcFibonacciValue(n, fibNumber)
}

func CalcFibonacciValue(n int64, fibCache FibCache) int64 {
	if _, ok := fibCache[n]; ok {
		return fibCache[n]
	}
	if n <= 1 {
		fibCache[n] = n
		return n
	}

	fibCache[n] = CalcFibonacciValue(n-1, fibCache) + CalcFibonacciValue(n-2, fibCache)
	return fibCache[n]
}

func FibCalculator() func(int64) int64 {
	var fibNumber FibCache
	return func(n int64) int64 {
		fibNumber = make(FibCache)
		return CalcFibonacciValue(n, fibNumber)
	}
}
