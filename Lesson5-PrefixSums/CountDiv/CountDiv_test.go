package CountDiv

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A int
		B int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{6, 11, 2}, 3},
		{"simple", args{11, 345, 17}, 20},
		{"minimal", args{0, 0, 11}, 1},
		{"minimal", args{1, 1, 11}, 0},
		{"minimal", args{0, 1, 11}, 1},
		{"minimal", args{1, 0, 11}, 0},
		{"extreme_ifempty", args{10, 10, 5}, 1},
		{"extreme_ifempty", args{10, 10, 7}, 0},
		{"extreme_ifempty", args{10, 10, 20}, 0},
		//{"extreme_endpoint", args{?, ?, ?}, ?},		verify handling of range endpoints, multiple runs
		{"big_value1", args{100, 123000000, 2}, 61499951},
		{"big_value2", args{101, 123000000, 10000}, 12300},
		//{"big_value3", args{0, MAXINT, IN{1,MAXINT}}, ?},
		//{"big_value4", args{IN{1,MAXINT}, IN{1,MAXINT}, IN{1,MAXINT}}, ?},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A, tt.args.B, tt.args.K); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
