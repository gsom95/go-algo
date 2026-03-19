package stack

// Interface defines the methods that a stack data structure should implement.
type Interface[T any] interface {
	// Push adds an element to the top of the stack.
	Push(value T)

	// Pop removes and returns the element at the top of the stack.
	// If the stack is empty, it returns the zero value of T and false.
	Pop() (value T, ok bool)

	// Peek returns the element at the top of the stack without removing it.
	// If the stack is empty, it returns the zero value of T and false.
	Peek() (value T, ok bool)

	// IsEmpty returns true if the stack has no elements, otherwise false.
	IsEmpty() bool

	// Size returns the number of elements currently in the stack.
	Size() int
}
