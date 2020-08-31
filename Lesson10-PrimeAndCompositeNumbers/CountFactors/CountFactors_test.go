package CountFactors

import (
	"testing"
	"time"
)

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{24}, 8},
		{"2", args{1}, 1},
		{"2", args{2}, 2},
		{"2", args{3}, 2},
		{"2", args{4}, 3},
		{"2", args{5}, 2},
		{"2", args{6}, 4},
		{"2", args{7}, 2},
		{"2", args{8}, 4},
		{"2", args{9}, 3},
		{"2", args{10}, 4},
		{"3", args{16}, 5},
		{"4", args{36}, 9},
		{"6", args{64}, 7},
		{"7", args{4999696}, 45},
		{"8", args{39992976}, 135},		// limit 0.100 sec.
		{"9", args{2147395600}, 135},	// limit 0.100 sec.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeout := time.After(100 * time.Millisecond)
			done := make(chan bool)
			go func() {
				if got := Solution(tt.args.N); got != tt.want {
					t.Errorf("Solution() = %v, want %v", got, tt.want)
				}
				done <- true
			}()
			select {
			case <- timeout:
				t.Fatal("test out of time limit")
			case <-done:

			}
		})
	}
}