package Flags

import "math"

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}

	peaks := []int{}

	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}
	if len(peaks) == 0 {
		return 0
	}

	maxPossibleFlagCnt := int(math.Sqrt(float64(peaks[len(peaks)-1] - peaks[0]))) + 1

	for flag := maxPossibleFlagCnt; flag > 0; flag -= 1 {
		flagCnt := 1
		tmp := 0
		for j := 1; j < len(peaks); j++ {
			tmp += peaks[j] - peaks[j-1]
			if tmp >= flag {
				flagCnt += 1
				tmp = 0
			}
		}
		if flagCnt >= flag {
			return flag
		}
	}
	return 0
}
