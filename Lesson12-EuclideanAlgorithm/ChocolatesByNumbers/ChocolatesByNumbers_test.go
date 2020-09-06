package ChocolatesByNumbers

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N int
		M int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{10, 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N, tt.args.M); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}