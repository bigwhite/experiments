package stack

import "testing"

func TestStack(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	if s.Len() != 2 {
		t.Errorf("want 2, actual %d", s.Len())
	}

	top := s.Top()
	if top != 2 {
		t.Errorf("want 2, actual %d", top)
	}

	t1 := s.Pop()
	if t1 != 2 {
		t.Errorf("want 2, actual %d", t1)
	}
	t2 := s.Pop()
	if t2 != 1 {
		t.Errorf("want 1, actual %d", t2)
	}
}
