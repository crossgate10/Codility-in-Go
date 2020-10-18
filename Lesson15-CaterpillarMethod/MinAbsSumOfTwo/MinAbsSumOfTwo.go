package MinAbsSumOfTwo

import (
	"sort"
)

func Solution(A []int) int {
	diff := int(20E8)
	start := 0
	end := len(A) - 1

	sort.Ints(A)
	for start <= end {
		diff = min(diff, abs(A[start] + A[end]))

		if abs(A[start]) > abs(A[end]) {
			start += 1
		} else {
			end -= 1
		}
	}
	return diff
}

func min(a, b int) int  {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}