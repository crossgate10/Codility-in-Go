package FibonacciNumbers

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"1", args{7}, 13},
		{"2", args{10}, 55},
		{"3", args{42}, 267914296},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib1(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFib1_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib1(10)
	}
}

func BenchmarkFib1_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib1(20)
	}
}

func TestFib2(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"1", args{7}, 13},
		{"2", args{10}, 55},
		{"3", args{42}, 267914296},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib2(tt.args.n); got != tt.want {
				t.Errorf("Fib2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFib2_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib2(10)
	}
}

func BenchmarkFib2_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib2(20)
	}
}

func TestFibBig(t *testing.T) {
	l1, _ := new(big.Int).SetString("354224848179261915075", 10)
	l2, _ := new(big.Int).SetString("176023680645013966468226945392411250770384383304492191886725992896575345044216019675", 10)
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"1", args{7}, big.NewInt(13)},
		{"2", args{10}, big.NewInt(55)},
		{"3", args{42}, big.NewInt(267914296)},
		{"4", args{100}, l1},
		{"5", args{400}, l2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibBig(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FibBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibBig_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibBig(10)
	}
}

func BenchmarkFibBig_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibBig(20)
	}
}