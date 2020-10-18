package MissingInteger

import "sort"

func Solution(A []int) int {
	missing := 1
	sort.Ints(A)
	for _, e := range A {
		if e == missing {
			missing++
		} else if e > missing {
			break
		}
	}
	return missing
}