package FrogRiverOne

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type Input struct {
		X int
		A []int
	}
	type Case struct {
		Input    Input
		Expected int
		Actual   int
	}

	ns := []Case{
		{Input{5, []int{1,3,1,4,2,3,5,4}}, 6, 0},
	}

	for _, n := range ns {
		//n := n
		t.Run(fmt.Sprintf("Case-%v", n.Input), func(t *testing.T) {
			//t.Parallel()
			if actual := Solution(n.Input.X, n.Input.A); actual != n.Expected {
				t.Errorf("Expected: %v, actual: %v", n.Expected, actual)
			}
		})
	}
}

