package MaxSliceSum

import (
	"fmt"
	"math"
)

func Solution(A []int) int {
	leftSlice, untilNow, onceTotal := 0, 0, 0

	for i := 3; i < len(A); i++ {
		leftSlice = int(math.Max(0, float64(A[i-2]+leftSlice)))
		untilNow = int(math.Max(float64(leftSlice), float64(A[i-1])+float64(untilNow)))
		onceTotal = int(math.Max(float64(untilNow), float64(onceTotal)))
		fmt.Print(1)
	}
	return onceTotal
}
