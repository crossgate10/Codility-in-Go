package AbsDistinct

func Solution(A []int) int {
	keys := make(map[int]bool)
	list := []int{}
	for k := range A {
		A[k] = Abs(A[k])
		if _, ok := keys[A[k]]; !ok {
			keys[A[k]] = true
			list = append(list, A[k])
		}
	}

	return len(list)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
