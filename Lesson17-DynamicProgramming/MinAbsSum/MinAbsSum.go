package MinAbsSum

func Solution(A []int) int {
	n := len(A)
	m := 0

	for i := 0; i < n; i++ {
		A[i] = abs(A[i])
		m = max(A[i], m)
	}

	s := sum(A)
	cnt := make([]int, m+1)

	for i := 0; i < n; i++ {
		cnt[A[i]] += 1
	}

	dp := make([]int, s+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}

	for i := 1; i < m+1; i++ {
		if cnt[i] > 0 {
			for j := 0; j < s; j++ {
				if dp[j] >= 0 {
					dp[j] = cnt[i]
				} else if j >= i && dp[j - i] > 0 {
					dp[j] = dp[j-i] -1
				}
			}
		}
	}
	res := s

	for i := 0; i < (s/2+1); i++ {
		if dp[i] >= 0 {
			res = min(res, s - 2*i)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(a []int) int {
	res := 0
	for i := 0; i < len(a); i++ {
		res += a[i]
	}
	return res
}

