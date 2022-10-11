package golang

import "fmt"

type MyIntegerSet struct {
	set map[int]bool
}

func NewSet() MyIntegerSet {
	m := make(map[int]bool)
	return MyIntegerSet{m}
}

func (s *MyIntegerSet) Add(elem int) {
	if !s.IsPresent(elem) {
		s.set[elem] = true
	}
}

func (s *MyIntegerSet) IsPresent(elem int) bool {
	for k, v := range s.set { //the use of k,v is ok in golang. golang likes small variables
		if elem == k {
			fmt.Printf("element: %v is present: %v", elem, v)
			return true
		}
	}
	return false
}
