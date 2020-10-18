package PassingCars

func Solution(A []int) int {
	sums := prefixSums(A)
	cnt := 0

	for idx, val := range A {
		if val == 0 {
			cnt += countTotal(sums, idx, len(A)-1)
		}
	}

	if cnt > 1000000000 {
		return -1
	} else {
		return cnt
	}
}

func prefixSums(A []int) []int {
	n := len(A)
	p := make([]int, n +1)
	for i:= 1; i< n+1; i++ {
		p[i] = p[i - 1] + A[i - 1]
	}

	return p
}

func countTotal(p []int, x, y int) int {
	return p[y+1] - p[x]
}