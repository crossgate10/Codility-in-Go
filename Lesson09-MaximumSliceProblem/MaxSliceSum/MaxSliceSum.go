package MaxSliceSum

import "math"

func Solution(A []int) int {
	maxSum := float64(1 - (1 << 32 -1))
	maxEnding := maxSum

	for _, value := range A {
		maxEnding = math.Max(float64(value), float64(value) + maxEnding)
		maxSum = math.Max(maxEnding, maxSum)
	}
	return int(maxSum)
}
