package binary_gap

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type Case struct {
		Input    int
		Expected int
		Actual   int
	}
	ns := []Case{
		{32, 0, 0},
		{15, 0, 0},
		{1041, 5, 0},
	}

	for _, n := range ns {
		//n := n
		t.Run(fmt.Sprintf("Case-%d", n.Input), func(t *testing.T) {
			//t.Parallel()
			if actual := Solution(n.Input); actual != n.Expected {
				t.Errorf("Expected: %v, actual: %v", n.Expected, actual)
			}
		})
	}
}
