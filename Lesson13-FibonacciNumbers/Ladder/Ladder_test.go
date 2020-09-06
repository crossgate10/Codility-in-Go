package Ladder

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{4, 4, 5, 5, 1}, []int{3, 2, 4, 3, 1}}, []int{5, 1, 8, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
