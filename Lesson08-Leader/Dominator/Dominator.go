package Dominator

import (
	"codilityingo/pkg/StackQueue"
	"math"
)

func Solution(A []int) int {
	arrLen := len(A)

	if arrLen == 1 {
		return 0
	}

	l := StackQueue.NewIntStack(arrLen)
	candidate := -1
	leader := -1
	cnt := 0

	for _, value := range A {
		if l.Size() == 0 {
			l.Push(value)
		} else {
			if value != l.Front() {
				l.Pop()
			} else {
				l.Push(value)
			}
		}
	}

	if l.Size() > 0 {
		candidate = l.Front()
	} else {
		return -1
	}

	for _, value := range A {
		if value == candidate {
			cnt += 1
		}
	}

	if cnt > int(math.Floor(float64(arrLen)/2.0)) {
		leader = candidate
	} else {
		return -1
	}
	return FindIndex(arrLen, func(i int) bool {
		return A[i] == leader
	})
}

func FindIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
