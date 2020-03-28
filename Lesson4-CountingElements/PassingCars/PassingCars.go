package Temp

import "sort"

func Solution(A []int) int {
	sort.Ints(A)
	resp := 1
	for i, e := range A {
		if e == i+1 {
			continue
		} else {
			resp = 0
		}
	}
	return resp
}
