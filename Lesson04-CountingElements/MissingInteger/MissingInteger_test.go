package MissingInteger

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
		{Input{[]int{1,3,6,4,1,2}}, 5, 0},
		{Input{[]int{1,2,3}}, 4, 0},
		{Input{[]int{-1,-3}}, 1, 0},
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

