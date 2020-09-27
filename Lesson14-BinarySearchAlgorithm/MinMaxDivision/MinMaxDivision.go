package MinMaxDivision

func Solution(K, N int, A []int) int {
	lowerBound := Max(A)
	upperBound := Sum(A)

	if K == 1 {
		return upperBound
	}
	if K >= len(A) {
		return lowerBound
	}

	for lowerBound <= upperBound {
		candidate := (lowerBound + upperBound) / 2

		if check(A, K, candidate) {
			upperBound = candidate - 1
		} else {
			lowerBound = candidate + 1
		}
	}
	return lowerBound
}

func check(a []int, k, maxBlock int) bool {
	blockSum := 0
	cnt := 0

	for _, v := range a {
		v := v
		if blockSum+v > maxBlock {
			blockSum = v
			cnt += 1
		} else {
			blockSum += v
		}

		if cnt >= k {
			return false
		}
	}
	return true
}

func Max(a []int) int {
	max := -999
	for _, v := range a {
		v := v
		if v > max {
			max = v
		}
	}
	return max
}

func Sum(a []int) int {
	res := 0
	for i := 0; i < len(a); i++ {
		res += a[i]
	}
	return res
}
