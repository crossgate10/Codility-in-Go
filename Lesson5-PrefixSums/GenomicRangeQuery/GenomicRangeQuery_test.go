package GenomicRangeQuery

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		S string
		P []int
		Q []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"first", args{"AC", []int{0, 0, 1}, []int{0, 1, 1}}, []int{1, 1, 2}},
		{"second", args{"A", []int{0}, []int{0}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S, tt.args.P, tt.args.Q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}