package CyclicRotation

func Solution(A []int, K int) []int {
	// Score:100%
	if K < 0 || len(A) == 0 {
		return A
	}
	r := len(A) - K%len(A)
	A = append(A[r:], A[:r]...)
	return A
}