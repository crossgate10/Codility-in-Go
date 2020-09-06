package EuclideanAlgorithm

import "testing"

func TestGCD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{15, 75}, 15},
		{"2", args{10, 15}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	type args struct {
		a        int
		b        int
		integers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{10, 15, []int{}}, 30},
		{"2", args{10, 15, []int{20}}, 60},
		{"3", args{1, 2, []int{3, 4, 5, 6, 7, 8, 9, 10}}, 2520},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.args.a, tt.args.b, tt.args.integers...); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}
