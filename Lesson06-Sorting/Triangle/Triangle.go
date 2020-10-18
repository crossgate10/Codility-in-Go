package Triangle

import "sort"

func Solution(A []int) int {
	l := len(A)
	res := 0

	if l < 3 {
		return res
	}

	sort.Ints(A)

	for i := 0; i < l-2; i += 1 {
		if A[i+2] < A[i+1]+A[i] {
			res = 1
			break
		}
	}
	return res
}
