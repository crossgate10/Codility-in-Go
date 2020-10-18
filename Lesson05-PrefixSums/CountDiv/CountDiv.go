package CountDiv

import (
	"math"
)

func Solution(A int, B int, K int) int {
	addOne := 0
	if A%K == 0 {
		addOne = 1
	}
	return int(math.Floor(float64(B/K))) - int(math.Floor(float64(A/K))) + addOne
}
