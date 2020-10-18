package CountTriangles

import "sort"

func Solution(A []int) int {
	n := len(A)
	res := 0

	sort.Ints(A)
	for x := 0; x < n; x++ {
		z := x + 2
		for y := x + 1; y < n; y++ {
			for z < n && A[x]+A[y] > A[z] {
				z += 1
			}
			res += z - y - 1
		}
	}
	return res
}
