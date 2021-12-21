package main_test

import (
	. "gb/lesson10/lesson4/sort"
	"math/rand"
	"sort"
	"testing"
)

func generateRandomSlice(n int) []int {
	targetSlice := []int{}
	for i := 1; i < n; i++ {
		randI := rand.Intn(i)
		targetSlice = append(targetSlice, randI)
	}

	return targetSlice
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestCalcFibonacciValueFull(t *testing.T) {
	for i := 1; i < 100; i++ {
		targetSlice := generateRandomSlice(i)
		copyOfTargetSlice := targetSlice
		Sort(targetSlice)
		sort.Ints(copyOfTargetSlice)
		if !Equal(targetSlice, copyOfTargetSlice) {
			t.Errorf("Result of sorting is not valid on %d iteration, expected %d, but received %d", i, targetSlice, copyOfTargetSlice)
		}
	}
}

func BenchmarkSort(b *testing.B) {
	targetSlice := generateRandomSlice(100)
	for i := 0; i < b.N; i++ {
		Sort(targetSlice)
	}
}
