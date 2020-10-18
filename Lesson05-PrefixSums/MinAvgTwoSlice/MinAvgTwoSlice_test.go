package MinAvgTwoSlice

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first", args{A: []int{1, 3, 6, 4, 1, 2}}, 4},
		{"second", args{A: []int{10, 10, -1, 2, 4, -1, 2, -1}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
