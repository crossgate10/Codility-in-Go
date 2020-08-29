package NumberOfDiscIntersections

import "sort"

func Solution(A []int) int {
	lefts := make([]int, len(A))  // 每個圓最左側的位置
	rights := make([]int, len(A)) // 每個圓最右側的位置
	res := 0

	for i := 0; i < len(A); i += 1 {
		lefts[i] = i - A[i]
		rights[i] = i + A[i]
	}

	sort.Ints(lefts)
	sort.Ints(rights)

	for i := 0; i < len(A); i++ {
		end := rights[i]

		count := sort.Search(len(lefts), func(i int) bool {
			return lefts[i] > end
		})

		count -= (i + 1)
		res += count

		if res > 10000000 {
			return -1
		}
	}
	return res
}
