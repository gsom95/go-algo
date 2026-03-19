package stack

import "testing"

func TestListStackPush(t *testing.T) {
	s := ListStack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
}

func TestListStackPop(t *testing.T) {
	s := ListStack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	value, ok := s.Pop()
	if !ok {
		t.Errorf("Expected pop to succeed, but it failed")
	}
	if value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}
}

func TestListStackPeek(t *testing.T) {
	s := ListStack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	value, ok := s.Peek()
	if !ok {
		t.Errorf("Expected peek to succeed, but it failed")
	}
	if value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
}

func TestListStackIsEmpty(t *testing.T) {
	s := ListStack[int]{}
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty, but it is not")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("Expected stack to not be empty, but it is")
	}
}

func TestListStackSize(t *testing.T) {
	s := ListStack[int]{}
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	s.Push(1)
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}
}
