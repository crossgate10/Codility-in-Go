package MaxProfit

import (
	"math"
)

func Solution(A []int) int {
	minPrice := float64(1 << 32 -1)
	maxProfit := 0.0

	for _, val := range A {
		minPrice = math.Min(minPrice, float64(val))
		maxProfit = math.Max(float64(val) - minPrice, maxProfit)
	}
	return int(maxProfit)
}
