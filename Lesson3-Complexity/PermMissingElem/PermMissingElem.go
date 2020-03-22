package PermMissingElem

func Solution(A []int) int {
	l := len(A)+1
	total := (1+l)*l/2
	for i := range A {
		total -= A[i]
	}
	return total
}