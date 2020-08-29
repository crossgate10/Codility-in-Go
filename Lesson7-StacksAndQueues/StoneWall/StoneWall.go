package StoneWall

import "codilityingo/Lesson7-StacksAndQueues/StackQueue"

func Solution(H []int) int {
	l := StackQueue.NewIntStack(len(H))
	res := 0

	for _, value := range H {
		for l.Size() > 0 && l.Front() > value {
			l.Pop()
		}

		if l.Size() == 0 || l.Front() < value {
			l.Push(value)
			res += 1
		}
	}
	return res
}
