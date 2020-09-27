package NailingPlanks

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		B []int
		C []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}