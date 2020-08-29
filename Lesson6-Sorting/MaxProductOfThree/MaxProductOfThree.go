package MaxProductOfTree

import "sort"

func Solution(A []int) int {
	l := len(A)

	sort.Ints(A)
	o1 := A[l-1] * A[l-2] * A[l-3]
	o2 := A[l-1] * A[0] * A[1]
	if o1 > o2 {
		return o1
	} else {
		return o2
	}
}
