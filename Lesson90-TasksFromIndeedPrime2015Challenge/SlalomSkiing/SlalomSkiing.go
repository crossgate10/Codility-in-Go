package SlalomSkiing

func Solution(A []int) int {
	bound := max(A) + 1
	multiverse := []int{}
	for i := 0; i < len(A); i++ {
		multiverse = append(multiverse, bound * 2 + A[i])
		multiverse = append(multiverse, bound * 2 - A[i])
		multiverse = append(multiverse, A[i])
	}
	return lis(multiverse)
}

func lis(A []int) int {
	smallestEndVal := make([]int, len(A)+1)
	smallestEndVal[0] = -1
	lisLen := 0
	for i := 0; i < len(A); i++ {
		lower := 0
		upper := lisLen
		for lower <= upper {
			mid := (upper + lower) / 2
			if A[i] < smallestEndVal[mid] {
				upper = mid - 1
			} else if A[i] > smallestEndVal[mid] {
				lower = mid + 1
			} else {
				// happen when element of A not distinct
			}
		}
		if smallestEndVal[lower] == 0 {
			smallestEndVal[lower] = A[i]
			lisLen += 1
		} else {
			smallestEndVal[lower] = min(smallestEndVal[lower], A[i])
		}
	}
	return lisLen
}

func max(A []int) int {
	maxElement := -1
	for _, e := range A {
		e := e
		if e > maxElement {
			maxElement = e
		}
	}
	return maxElement
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
