package CountSemiprimes

func Solution(N int, P, Q []int) []int {
	primeTable := make([]bool, N+1)
	for i := 0; i < N+1; i++ {
		if i != 0 && i != 1 {
			primeTable[i] = true
		}
	}
	prime := []int{}
	semiPrime := make([]int, N+1)
	result := []int{}

	for idx := 2; idx*idx <= N; idx++ {
		for i := 2; idx*i <= N; i++ {
			primeTable[idx*i] = false
		}
	}
	for i := 0; i < len(primeTable); i++ {
		if primeTable[i] {
			prime = append(prime, i)
		}
	}
	// 計算半質數陣列
	for i := 0; i < len(prime); i++ {
		for j := i; j < len(prime); j++ {
			tmp := prime[i] * prime[j]
			if tmp <= N {
				semiPrime[tmp] = 1
			} else {
				break
			}
		}
	}
	// prefix sum algorithm
	for i := 1; i < N+1; i++ {
		semiPrime[i] += semiPrime[i-1]
	}
	for i := 0; i < len(P); i++ {
		result = append(result, semiPrime[Q[i]]-semiPrime[P[i]-1])
	}

	return result
}
