package semantic

type Stack[T any] []T

func (s *Stack[T]) Top() (t T) {
	l := len(*s)
	if l == 0 {
		return t
	}
	return (*s)[l-1]
}

func (s *Stack[T]) Push(v T) {
	// for debug
	(*s) = append((*s), v)

	//fmt.Printf("push %v:%v, stack len[%d]\n", v.Type(), v.Value(), len(*s))
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Pop() T {
	if len(*s) < 1 {
		panic("can not pop because stack is empty")
	}

	// Get the last element from the stack.
	result := (*s)[len(*s)-1]

	// Remove the last element from the stack.
	*s = (*s)[:len(*s)-1]

	// for debug
	//fmt.Printf("pop %v:%v, stack len[%d]\n", result.Type(), result.Value(), len(*s))

	return result
}
