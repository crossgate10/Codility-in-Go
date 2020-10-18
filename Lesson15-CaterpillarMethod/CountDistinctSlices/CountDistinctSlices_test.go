package CountDistinctSlices

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		M int
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"debug", args{6, []int{3, 4, 5, 5, 2}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.M, tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
