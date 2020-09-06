package FibFrog

func Solution(A []int) int {
	// 添加開始與結束的位置
	A = prependInt(A, 1)
	A = append(A, 1)

	l := len(A)
	fibList := fib(l)
	if l-1 == fibList[len(fibList) - 1] {
		// 一次完成
		return 1
	} else {
		signList := make([]int, l)
		for i := 1; i < l; i++ {
			signList[i] = l
			if A[i] == 1 {
				// 遍歷翡波那契數列，尋找最少跳躍次數
				for j := 0; j < len(fibList); j++ {
					dis := i - fibList[j]
					if dis >= 0 {
						// dis 這個位置可用翡波那契數到達
						if signList[dis] + 1 < signList[i] {
							signList[i] = signList[dis] + 1		// 達到位置dis的次數 + 走長度為翡波那契數j的一次
						}
					} else {
						break	// 表示翡波那契數值更大，則跳出
					}
				}
			}
		}
		if signList[len(signList)-1] < l {
			return signList[len(signList)-1]
		} else {
			return -1
		}
	}
}

func prependInt(A []int, x int) []int {
	A = append(A, 0)
	copy(A[1:], A)
	A[0] = x
	return A
}

func fib(n int) []int {
	res := []int{1}
	a, b := 1, 1
	for b <= n {
		res = append(res, b)
		a, b = b, a+b
	}
	return res
}
