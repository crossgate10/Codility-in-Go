package TieRopes

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		K int
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"debug", args{
			K: 4,
			A: []int{1,2,3,4,1,1,3},
		}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.K, tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}