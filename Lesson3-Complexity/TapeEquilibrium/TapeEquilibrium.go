package TapeEquilibrium

func Solution(A []int) int {
	return Solution1(A)
}

func Solution1(A []int) int {
	l, diff, left, right, absSum, min := len(A), 0, 0, 0, 0, 0
	for i := 0; i < l; i++ {
		right += A[i]
		absSum += abs(A[i])
	}
	for i := 0; i < l-1; i++ {
		right -= A[i]
		left += A[i]
		diff = abs(left - right)
		if i == 0 {
			min = diff
		}
		if diff < min {
			min = diff
		}
	}
	return min
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
