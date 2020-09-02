package Peaks

func Solution(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}

	peaks := []int{}
	for i := 1; i < n-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	for block := 3; block < (n/2)+1; block++ {
		if n%block == 0 {
			idxToCheck := 0
			tmpBlocks := 0
			for _, peak := range peaks {
				if peak >= idxToCheck && peak < idxToCheck + block {
					tmpBlocks += 1
					idxToCheck += block
				}
			}
			if n/block == tmpBlocks {
				return tmpBlocks
			}
		}
	}
	if len(peaks) > 0 {
		return 1
	}
	return 0
}
