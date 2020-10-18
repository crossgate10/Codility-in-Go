package TapeEquilibrium

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type Input struct {
		A []int
	}
	type Case struct {
		Input    Input
		Expected int
		Actual   int
	}

	ns := []Case{
		{Input{[]int{3, 1, 2, 4, 3}}, 1, 0},
		{Input{[]int{-1000, 1000}}, 2000, 0},
		{Input{[]int{-1000, -1000, 1000, 1000}}, 2000, 0},
		{Input{[]int{-2, -3, -4, -1}}, 0, 0},
		{Input{[]int{-10, -20, -30, -40, 100}}, 20, 0},
		{Input{[]int{-10, -40, -50, 100}}, 20, 0},
	}

	for _, n := range ns {
		//n := n
		t.Run(fmt.Sprintf("Case-%v", n.Input), func(t *testing.T) {
			//t.Parallel()
			if actual := Solution(n.Input.A); actual != n.Expected {
				t.Errorf("Expected: %v, actual: %v", n.Expected, actual)
			}
		})
	}
}
