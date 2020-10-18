package MinAvgTwoSlice

func Solution(A []int) int {
    arraylen := len(A)
    min := float64(1 << 32 -1)
    result := 0

    for i:= 0; i < arraylen; i++ {
        if i+1 < arraylen {
            temp := (float64(A[i]) + float64(A[i+1]))/float64(2.0)

            if temp < float64(min) {
                min = temp
                result = i
            }
        }

        if i+2 < arraylen {
            temp := (float64(A[i]) + float64(A[i+1]) + float64(A[i+2]))/float64(3.0)

            if temp < float64(min) {
                min = temp
                result = i
            }
        }
    }
    return result
}
