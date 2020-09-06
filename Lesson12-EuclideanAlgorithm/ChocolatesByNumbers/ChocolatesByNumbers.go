package ChocolatesByNumbers

func Solution(N int, M int) int {
	totalChoco := N
	for M != 0 {
		t := M
		M = N % M
		N = t
	}
	return totalChoco / N
}
