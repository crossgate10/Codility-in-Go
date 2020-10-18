package FrogRiverOne

func Solution(X int, A []int) int {
	return Solution2(A, X)
	//return Solution1(X, A)
}

func Solution2(A []int, D int) int {
	l := len(A)
	tiles := make([]bool, l)
	todo := D

	for i:=0; i<l; i++ {
		inIdx := A[i] - 1
		if !tiles[inIdx] {
			todo--
			tiles[inIdx] = true
		}
		if todo == 0 {
			return i
		}
	}
	return -1
}

func Solution1(X int, A []int) int {
	s := Set{map[int]Empty{}}
	for i:=0; i< len(A); i++ {
		s.Add(A[i])
		if s.Size() == X {
			return i
		}
	}
	return -1
}

var empty Empty

type Empty struct {}

type Set struct {
	m map[int]Empty
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Add(n int) {
	s.m[n] = empty
}