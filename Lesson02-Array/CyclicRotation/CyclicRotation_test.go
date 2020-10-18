package CyclicRotation

import (
	"fmt"
	"testing"
	"helper"
)

func TestSolution(t *testing.T) {
	type Input struct {
		A []int
		K int
	}
	type Case struct {
		Input    Input
		Expected []int
		Actual   []int
	}

	ns := []Case{
		{Input{[]int{3,8,9,7,6}, 3}, []int{9,7,6,3,8}, []int{}},
		{Input{[]int{0,0,0}, 1}, []int{0,0,0}, []int{}},
		{Input{[]int{1,2,3,4}, 4}, []int{1,2,3,4}, []int{}},
	}

	for _, n := range ns {
		//n := n
		t.Run(fmt.Sprintf("Case-%v", n.Input), func(t *testing.T) {
			//t.Parallel()
			if actual := Solution(n.Input.A, n.Input.K); !Equal(actual, n.Expected) {
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