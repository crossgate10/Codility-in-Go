package MaxCounters

func Solution(N int, A []int) []int {
	l, nmax, amax, ans := len(A), 0, 0, make([]int,N)
	for i:=0; i<l; i++ {
		if A[i] <= N {
			if ans[A[i]-1]<nmax {
				ans[A[i]-1] = nmax + 1
			}else{
				ans[A[i]-1]++
			}
			amax = max(amax, ans[A[i]-1])	// 每回合迭代的最大值
		}else{
			nmax = amax		// 將當前最大值指派給n次迭代的最大值，作為閥值
		}
		// fmt.Println(A[i], nmax, amax, ans)
	}

	// 最後只要值小於閥值，則 = n次迭代最大值
	for i:=0; i<N; i++ {
		if ans[i]<nmax {
			ans[i] = nmax
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
