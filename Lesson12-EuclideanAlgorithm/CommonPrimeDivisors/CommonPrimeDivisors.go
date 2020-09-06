package CommonPrimeDivisors

func Solution(A, B []int) int {
	cnt := 0
	for i := 0; i < len(B); i++ {
		a, b := A[i], B[i]
		x := gcd(a, b)

		for true {
			d := gcd(x, a)
			if d == 1 {
				break
			}
			a /= d
		}

		for true {
			d := gcd(x, b)
			if d == 1 {
				break
			}
			b /= d
		}
		if a == 1 && b == 1 {
			cnt += 1
		}
	}
	return cnt
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}