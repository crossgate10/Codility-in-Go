package Temp

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
		{Input{[]int{4,1,3,2}}, 1, 0},
		{Input{[]int{1,4,3}}, 0, 0},
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

