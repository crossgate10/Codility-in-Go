package MaxNonoverlappingSegments

func Solution(A, B []int) int {
	n := len(A)
	if n < 1 {
		return 0
	}

	cnt := 1
	prevEnd := B[0]
	for i := 0; i < n; i++ {
		if A[i] > prevEnd {
			cnt += 1
			prevEnd = B[i]
		}
	}
	return cnt
}
