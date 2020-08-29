package Distinct

import "sort"

func Solution(A []int) int {
	//return Solution1(A)
	return Solution2(A)
}

func Solution1(A []int) int {
	// 粗暴簡單
	result := []int{}
	for i := range A {
		exist := false
		for j := range result {
			if A[i] == result[j] {
				exist = true
			}
		}
		if !exist {
			result = append(result, A[i])
		}
	}
	return len(result)
}

func Solution2(A []int) int {
	arrayLen := len(A)
	sort.Ints(A)
	result := 1

	if arrayLen == 0 {
		return 0
	}

	for i := 1; i < arrayLen; i++ {
		if A[i] != A[i-1] {
			result += 1
		}
	}
	return result
}
