package MinMaxDivision

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		K int
		N int
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, 5, []int{2, 1, 5, 1, 2, 2, 2}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.K, tt.args.N, tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
