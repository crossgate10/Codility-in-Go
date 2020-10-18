package MonckyCrossRiver

import (
	"sort"
)

func Solution(A []int, D int) int {
	N := len(A)

	sort.Ints(A)

	at := 0
	check := 0
	best := -1
	jumps := 0
	for at < N {
		for check < N && A[check] <= at+D {
			best = check
			check++
		}

		if best >= 0 {
			at = A[best]
			jumps++
			//fmt.Printf("at: %d, check: %d, best: %d, jumps: %d\n", at, check, best, jumps)
			best = -1
		} else {
			break
		}
	}
	if at+D <= N {
		return -1
	}
	return jumps
}
