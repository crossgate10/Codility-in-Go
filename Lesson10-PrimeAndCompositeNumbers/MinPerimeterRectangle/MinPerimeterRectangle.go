package MinPerimeterRectangle

import "math"

func Solution(N int) int {
	for i := int(math.Sqrt(float64(N))); i > 0; i-- {
		if N%i == 0 {
			return (int(N/i) + i) * 2
		}
	}
	return 0
}
