package CountFactors

func Solution(N int) int {
	return Solution2(N)
}

func Solution1(N int) int {
	// Score: 71%
	res := []int{}
	for i := N; i > 0; i-- {
		if (N/i)*i == N {
			res = append(res, i)
		}
	}
	return len(res)
}

func Solution2(N int) int {
	// Score: 100%
	n := N
	i, cnt := 1, 0
	for ; i*i < n; i++ {
		if n % i == 0 {
			cnt += 2
		}
	}
	if i*i == n {
		cnt++
	}
	return cnt
}
