package stack

// ArrayStack is a stack implementation using a slice as the underlying data structure.
// Zero value of ArrayStack is an empty stack and can be used.
type ArrayStack[T any] struct {
	elements []T
}

var _ Interface[any] = (*ArrayStack[any])(nil)

func (s *ArrayStack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *ArrayStack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return element, true
}

func (s *ArrayStack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *ArrayStack[T]) Size() int {
	return len(s.elements)
}
