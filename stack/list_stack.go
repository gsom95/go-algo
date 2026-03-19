package stack

// ListStackElement represents an element in the linked list stack.
type ListStackElement[T any] struct {
	value T
	next  *ListStackElement[T]
}

// ListStack is a stack implementation using a linked list.
// It provides O(1) time complexity for push, pop, and peek operations.
type ListStack[T any] struct {
	top  *ListStackElement[T]
	size int
}

var _ Interface[any] = (*ListStack[any])(nil)

func (s *ListStack[T]) Push(value T) {
	s.top = &ListStackElement[T]{
		value: value,
		next:  s.top,
	}
	s.size++
}

func (s *ListStack[T]) Pop() (value T, ok bool) {
	if s.top == nil {
		var zero T
		return zero, false
	}
	value = s.top.value
	s.top = s.top.next
	s.size--
	return value, true
}

func (s *ListStack[T]) Peek() (value T, ok bool) {
	if s.top == nil {
		var zero T
		return zero, false
	}
	return s.top.value, true
}

func (s *ListStack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *ListStack[T]) Size() int {
	return s.size
}
