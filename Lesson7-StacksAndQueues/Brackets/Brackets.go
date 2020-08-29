package Brackets

import "codilityingo/Lesson7-StacksAndQueues/StackQueue"

func Solution(S string) int {
	l := StackQueue.NewStringStack(len(S))
	pairs := make(map[string]string)
	pairs["{"] = "}"
	pairs["["] = "]"
	pairs["("] = ")"

	if len(S) == 0 {
		return 1
	}
	if len(S) % 2 != 0 {
		return 0
	}

	for _, value := range S {
		val := string(value)
		if val == "(" ||val == "{" || val == "[" {
			l.Push(val)
		} else if val == ")" ||val == "}" || val == "]" {
			if l.Size() == 0 {
				return 0
			}

			head := l.Front()
			if pairs[head] == val {
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
