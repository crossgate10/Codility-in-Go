package EquiLeader

import (
	"codilityingo/pkg/StackQueue"
	"math"
)

func Solution(A []int) int {
	arrLen := len(A)
	leader, cnt := FindLeader(A, arrLen)
	if arrLen == 1 || leader == -1 {
		return 0
	}

	leadersCnt := 0
	leftLeadersCnt, leftSeqLen, rightSeqLen := 0, 0, 0
	for i, value := range A {
		if value == leader {
			leftLeadersCnt += 1
		}

		leftSeqLen = i + 1

		if leftLeadersCnt > int(math.Floor(float64(leftSeqLen) / 2.0)) {
			rightSeqLen = arrLen - leftSeqLen
			if cnt - leftLeadersCnt > int(math.Floor(float64(rightSeqLen) / 2.0)) {
				leadersCnt += 1
			}
		}
	}
	return leadersCnt
}

func FindLeader(A []int, arrLen int) (leader, cnt int) {
	l := StackQueue.NewIntStack(arrLen)
	candidate := -1
	leader = -1
	cnt = 0

	if arrLen == 1 {
		return
	}

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
		return
	}

	for _, value := range A {
		if value == candidate {
			cnt += 1
		}
	}

	if cnt > int(math.Floor(float64(arrLen) / 2.0)) {
		leader = candidate
	} else {
		return
	}
	return
}
