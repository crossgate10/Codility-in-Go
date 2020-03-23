package FrogRiverOne

func Solution(X int, A []int) int {
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