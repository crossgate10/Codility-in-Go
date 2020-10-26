package LongestPassword

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"test 5 a0A pass007 ?xy1"}, 7},
		{"2", args{"a b 11 2?3"}, -1},
		{"3", args{"pass0, logestPass001 111aa c 0"}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}