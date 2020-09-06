package SieveOfEratosthenes

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{30}, []int{2,3,5,7,11,13,17,19,23,29}},
		{"2", args{10}, []int{2,3,5,7}},
		{"3", args{60}, []int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sieve(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sieve() = %v, want %v", got, tt.want)
			}
		})
	}
}