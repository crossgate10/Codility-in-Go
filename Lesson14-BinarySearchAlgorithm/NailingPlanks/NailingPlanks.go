package NailingPlanks

import (
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

func Solution(A, B, C []int) int {
	return solution1(A, B, C)
	//return solution2(A, B, C)
}

func solution1(A, B, C []int) int {
	m := len(C)
	M := (m << 1) | 1
	left, right, res := 0, m, -1

	for left <= right {
		mid := (left + right) >> 1
		v := make([]int, M)
		for i := 0; i < mid; i++ {
			v[C[i]]++
		}
		for i := 1; i < M; i++ {
			v[i] += v[i-1]
		}
		can := true
		for i := 0; i < len(A); i++ {
			if v[B[i]]-v[A[i]-1] == 0 {
				can = false
				break
			}
		}
		if can {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func solution2(A, B, C []int) int {
	m := len(C)
	M := (m << 1) | 1
	const inf int = 2000000000
	nail := make([]int, M)
	for i := 0; i < M; i++ {
		nail[i] = inf
	}
	for i := m - 1; i >= 0; i-- {
		nail[C[i]] = i
	}

	plank := make([]int, M)
	for i := 0; i < len(A); i++ {
		plank[B[i]] = Max(plank[B[i]], A[i])
	}

	var left, right, r int
	q := deque.New()
	for i := 1; i < M; i++ {
		if plank[i] > left {
			left = plank[i]
			for !q.Empty() && (q.Left().(int) < left) {
				q.PopLeft()
			}
			for right = Max(right, left); right <= i; right++ {
				for !q.Empty() && nail[q.Right().(int)] >= nail[right] {
					q.PopRight()
				}
				q.PushRight(right)
			}
			r = Max(r, nail[q.Left().(int)])
			if r >= inf {
				return -1
			}
		}
	}
	return r + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
