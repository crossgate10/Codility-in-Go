package TieRopes

func Solution(K int, A []int) int {
	cnt, tmp := 0, 0

	for _, elem := range A {
		elem := elem
		tmp += elem
		if tmp >= K {
			cnt += 1
			tmp = 0
		}
	}
	return cnt
}
