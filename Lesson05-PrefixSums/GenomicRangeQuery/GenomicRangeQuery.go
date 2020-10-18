package GenomicRangeQuery

func Solution(S string, P []int, Q []int) []int {
	A, C, G := prefixSums(S)
	response := make([]int, len(P))

	for idx := range P {
		if A[Q[idx] + 1] - A[P[idx]] > 0 {
			response[idx] = 1
		} else if C[Q[idx] + 1] - C[P[idx]] > 0 {
			response[idx] = 2
		} else if G[Q[idx] + 1] - G[P[idx]] > 0 {
			response[idx] = 3
		} else {
			response[idx] = 4
		}
	}
	return response
}

func prefixSums(S string) ([]int, []int, []int) {
	n := len(S)
	A := make([]int, n+1)
	C := make([]int, n+1)
	G := make([]int, n+1)

	for i := 1; i < n + 1; i++ {
		s := string(S[i-1])
		A[i] = A[i-1]
		C[i] = C[i-1]
		G[i] = G[i-1]

		if s == "A" {
			A[i] += 1
		} else if s == "C" {
			C[i] += 1
		} else if s == "G" {
			G[i] += 1
		}
	}
	return A, C, G
}

