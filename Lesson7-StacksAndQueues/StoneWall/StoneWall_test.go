package StoneWall

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		H []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first", args{[]int{8,8,5,7,9,8,7,4,8}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.H); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}