package OddOccurrencesInArray

func Solution(A []int) int {
	// Score:100%
	result := 0
	for i := 0; i < len(A); i++ {
		result ^= A[i]
	}
	return result
}