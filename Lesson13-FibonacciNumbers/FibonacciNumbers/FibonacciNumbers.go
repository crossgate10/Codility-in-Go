package FibonacciNumbers

import "math/big"

func Fib1(n uint) uint {
	// Recursive approach
	if n <= 1 {
		return n
	}
	return Fib1(n - 1) + Fib1(n - 2)
}

func Fib2(n uint) uint64 {
	// Sequential approach
	if n <= 1 {
		return uint64(n)
	}

	var n2, n1 uint64 = 0,1

	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}
	return n2 + n1
}

func FibBig(n uint) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i:=uint(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}
	return n1
}
