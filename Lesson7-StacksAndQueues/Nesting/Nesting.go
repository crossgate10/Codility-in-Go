package Nesting

import "codilityingo/pkg/StackQueue"

func Solution(S string) int {
	l := StackQueue.NewStringStack(len(S))

	if len(S) == 0 {
		return 1
	}
	if len(S)%2 != 0 {
		return 0
	}

	for _, value := range S {
		val := string(value)
		if val == "(" {
			l.Push(val)
		} else if val == ")" {
			if l.Size() == 0 {
				return 0
			}

			head := l.Front()
			if head == "(" {
				l.Pop()
			} else {
				return 0
			}
		}
	}
	if l.Size() == 0 {
		return 1
	} else {
		return 0
	}
}
