package MonckyCrossRiver

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type Input struct {
		A []int
		D int
	}
	type Case struct {
		Input    Input
		Expected int
		Actual   int
	}

	ns := []Case{
		{Input{[]int{1, -1, 0, 2, 3, 5}, 3}, 2, 0},
		{Input{[]int{3, 2, 1}, 1}, 3, 0},
		{Input{[]int{1, 2, 3, 4, -1, -1, -1}, 3}, -1, 0},
	}

	for _, n := range ns {
		//n := n
		t.Run(fmt.Sprintf("Case-%v", n.Input), func(t *testing.T) {
			//t.Parallel()
			if actual := Solution(n.Input.A, n.Input.D); actual != n.Expected {
				t.Errorf("Expected: %v, actual: %v", n.Expected, actual)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	for i:=0; i< b.N; i++ {
		Solution([]int{1, 2, 3, 4, -1, -1, -1}, 3)
	}
}
