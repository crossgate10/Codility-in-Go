package FrogJump

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type Input struct {
		X int
		Y int
		D int
	}
	type Case struct {
		Input    Input
		Expected int
		Actual   int
	}

	ns := []Case{
		{Input{10,85,30}, 3, 0},
	}

	for _, n := range ns {
		//n := n
		t.Run(fmt.Sprintf("Case-%v", n.Input), func(t *testing.T) {
			//t.Parallel()
			if actual := Solution(n.Input.X, n.Input.Y, n.Input.D); actual != n.Expected {
				t.Errorf("Expected: %v, actual: %v", n.Expected, actual)
			}
		})
	}
}
