package binary_gap

import "strconv"
import "strings"

func Solution(N int) int {
	return Solution1(N)
	//return Solution2(N)

}

func Solution1(N int) int {
	// Score:60%

	n := int64(N)
	bits := strconv.FormatInt(n, 2)
	bitGroups := strings.Split(bits, "1")
	max := 0
	for i := range bitGroups {
		if len(bitGroups[i]) > max {
			max = len(bitGroups[i])
		}
	}

	if !(string(bits[0]) == "1" && string(bits[len(bits)-1]) == "1") {
		max = 0
	}
	return max
}

func Solution2(N int) int {
	// Score:100%

	var result int
	var result_tmp int
	var calc bool

	for N > 0 {
		if N%2 == 1 {
			if !calc {
				calc = true
			} else {
				if result_tmp > result {
					result = result_tmp
				}
				result_tmp = 0
			}
		} else {
			if calc {
				result_tmp++
			}
		}
		N /= 2
	}
	return result
}