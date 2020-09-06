package CountSemiprimes

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		N int
		P []int
		Q []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{26, []int{1, 4, 16}, []int{26, 10, 20}}, []int{10, 4, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N, tt.args.P, tt.args.Q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
