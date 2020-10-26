package FloodDepth

func Solution(A []int) int {
	left := calculateLeftMax(A)
	right := calculateRightMax(A)
	maxDepth := 0
	for i := 1; i < len(A)-1; i++ {
		leftMax := left[i]
		rightMax := right[i]
		h := A[i]
		if h < leftMax && h < rightMax {
			maxDepth = max(maxDepth, min(leftMax, rightMax)-h)
		}
	}
	return maxDepth
}

func calculateLeftMax(a []int) []int {
	n := len(a)
	left := make([]int, n)
	curMax := a[0]
	for i := 1; i < n; i++ {
		left[i] = curMax
		curMax = max(curMax, a[i])
	}
	return left
}

func calculateRightMax(a []int) []int {
	n := len(a)
	right := make([]int, n)
	curMax := a[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = curMax
		curMax = max(curMax, a[i])
	}
	return right
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
