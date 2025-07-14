package stack

type Stack[T any] struct {
	Items []T
}

func (s *Stack[T]) Push(element T) {
	if s.IsEmpty() {
		s.Items = make([]T, 0)
	}

	s.Items = append(s.Items, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.Items) - 1
	item := s.Items[index]
	s.Items = s.Items[:index]

	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	return s.Items[len(s.Items)-1], true
}

func (s *Stack[T]) Size() int {
	return len(s.Items)
}

func (s *Stack[T]) Clear() {
	s.Items = []T{}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Items) == 0
}
