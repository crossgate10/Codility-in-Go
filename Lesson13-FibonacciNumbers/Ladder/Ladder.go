package Ladder

import "math"

func Solution(A, B []int) []int {
	maxMum := max(A)
	l := len(A)
	fibTable := make([]int, maxMum+2)
	fibTable[0] = 0
	fibTable[1] = 1
	res := make([]int, l)

	// non-efficient fib
	for i := 2; i < maxMum+2; i++ {
		fibTable[i] = fibTable[i-2] + fibTable[i-1]
	}
	fibTable = fibTable[2:]

	for i := 0; i < l; i++ {
		val := fibTable[A[i]-1]
		res[i] = val & ((1 << B[i]) - 1)
	}
	return res
}

func max(A []int) int {
	max := -1
	for i := 0; i < len(A); i++ {
		max = int(math.Max(float64(A[i]), float64(max)))
	}
	return max
}
