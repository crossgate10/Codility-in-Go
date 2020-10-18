package CountDistinctSlices

import "log"

func Solution(M int, A []int) int {
	n := len(A)
	seen := make([]bool, M+1)
	front, back, count := 0, 0, 0

	log.Printf("front: %d, back:%d, count:%d seen:%v", front, back, count, seen)
	for front < n {
		if back < n && seen[A[back]] == false {
			seen[A[back]] = true
			back += 1
		} else {
			count += back - front
			seen[A[front]] = false
			front += 1
		}
		log.Printf("front: %d, back:%d, count:%d seen:%v", front, back, count, seen)
	}
	return Min(count, int(10e8))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
