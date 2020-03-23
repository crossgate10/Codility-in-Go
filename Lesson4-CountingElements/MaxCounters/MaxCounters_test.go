package MaxCounters

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type Input struct {
		N int
		A []int
	}
	type Case struct {
		Input    Input
		Expected []int
		Actual   []int
	}

	ns := []Case{
		{Input{5, []int{3,4,4,6,1,4,4}}, []int{3,2,2,4,2}, []int{}},
	}

	for _, n := range ns {
		//n := n
		t.Run(fmt.Sprintf("Case-%v", n.Input), func(t *testing.T) {
			//t.Parallel()
			if actual := Solution(n.Input.N, n.Input.A); !Equal(actual, n.Expected) {
				t.Errorf("Expected: %v, actual: %v", n.Expected, actual)
			}
		})
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
