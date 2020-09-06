package CountNonDivisible

func Solution(A []int) []int {
	// 紀錄已經算過的 A[i] 有多少除數
	toMapping := make(map[int]int)

	maxVal := 2 * len(A)
	occurs := make([]int, maxVal+1)
	for _, v := range A {
		occurs[v]++
	}

	r := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		totalDivisors := 0
		if _, ok := toMapping[A[i]]; ok {
			totalDivisors = toMapping[A[i]]
		} else {
			// 只需要計算到 √A[i]
			for j := 1; j*j <= A[i]; j++ {
				if j*j == A[i] {
					totalDivisors += occurs[j]
				} else {
					if A[i]%j == 0 {
						totalDivisors += occurs[j] + occurs[A[i]/j]		// 假設j是除數，A[i]/j也是除數
					}
				}
			}
			toMapping[A[i]] = totalDivisors
		}
		r[i] = len(A) - totalDivisors		// A的長度-除數數量 = 非除數數量
	}
	return r
}