package EuclideanAlgorithm

// GCD means Greatest Common Divisor
func GCD(a, b int) int {
	//return gcd1(a,b)
	return gcd2(a,b)
}

func gcd1(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func gcd2(a, b int) int {
	t := a % b
	if t > 0 {
		return gcd2(b, t)
	} else {
		return b
	}
}

// LCM means Least Common Multiple
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}


