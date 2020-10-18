package OddOccurrencesInArray

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
		{Input{[]int{9,3,9,3,9,7,9}}, 7, 0},
		{Input{[]int{0,0,0,1}}, 1, 0},
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
