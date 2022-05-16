package stack

type Stack[T any] []T

func (s *Stack[T]) Top() (t T) {
	l := len(*s)
	if l == 0 {
		return t
	}
	return (*s)[l-1]
}

func (s *Stack[T]) Push(v T) {
	(*s) = append((*s), v)
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Pop() (t T) {
	if len(*s) == 0 {
		return
	}

	// Get the last element from the stack.
	t = (*s)[len(*s)-1]

	// Remove the last element from the stack.
	*s = (*s)[:len(*s)-1]
	return
}
