package FrogJump

func Solution(X int, Y int, D int) int {
	result := ((Y-X)/D)
	remain := ((Y-X)%D)
	if remain > 0 {
		result++
	}
	return result
}