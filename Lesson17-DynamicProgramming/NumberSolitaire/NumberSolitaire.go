package NumberSolitaire

func Solution(A []int) int {
	n := len(A)
	dies := 6
	untilNow := make([]int, n+dies)

	for i := 0; i < len(untilNow); i++ {
		untilNow[i] = A[0]
	}

	for i := 1; i < n; i++ {
		untilNow[i+dies] = max(untilNow[i:i+dies]) + A[i]
	}
	return untilNow[len(untilNow)-1]
}

func max(a []int) int {
	m := - int(10E8)
	for i := 0; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}
