package Fish

import "codilityingo/pkg/StackQueue"

func Solution(A, B []int) int {
	l := StackQueue.NewIntStack(len(A))
	res := 0

	for i, _ := range B {
		if B[i] == 1 {
			l.Push(A[i])
		} else {
			for l.Size() > 0 {
				if l.Front() < A[i] {
					l.Pop()
				} else {
					break
				}
			}
			if l.Size() == 0 {
				res += 1
			}
		}
	}

	return res + l.Size()
}
