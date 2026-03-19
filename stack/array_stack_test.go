package stack

import (
	"slices"
	"testing"
)

func TestArrayStackPush(t *testing.T) {
	testCases := []struct {
		name      string
		initStack ArrayStack[int]
		input     []int
		expected  []int
	}{
		{
			name:      "push to empty stack",
			initStack: ArrayStack[int]{},
			input:     []int{1},
			expected:  []int{1},
		},
		{
			name:      "push to non-empty stack",
			initStack: ArrayStack[int]{elements: []int{4, 5, 6}},
			input:     []int{1, 2, 3},
			expected:  []int{4, 5, 6, 1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stack := tc.initStack
			for _, value := range tc.input {
				stack.Push(value)
			}
			if !slices.Equal(tc.expected, stack.elements) {
				t.Errorf("expected %v, got %v", tc.expected, stack.elements)
			}
		})
	}
}

func TestArrayStackPop(t *testing.T) {
	type popResulst struct {
		value int
		ok    bool
	}

	testCases := []struct {
		name                  string
		initStack             ArrayStack[int]
		testPop               func(*testing.T, *ArrayStack[int]) []popResulst
		expectedInStack       []int
		expectedPoppedResults []popResulst
	}{
		{
			name:      "pop from stack with one element",
			initStack: ArrayStack[int]{elements: []int{1}},
			testPop: func(t *testing.T, stack *ArrayStack[int]) []popResulst {
				result, ok := stack.Pop()
				return []popResulst{{value: result, ok: ok}}
			},
			expectedInStack:       []int{},
			expectedPoppedResults: []popResulst{{value: 1, ok: true}},
		},
		{
			name:      "pop from stack with multiple elements",
			initStack: ArrayStack[int]{elements: []int{1, 2, 3}},
			testPop: func(t *testing.T, as *ArrayStack[int]) []popResulst {
				result, ok := as.Pop()
				return []popResulst{{value: result, ok: ok}}

			},
			expectedInStack:       []int{1, 2},
			expectedPoppedResults: []popResulst{{value: 3, ok: true}},
		},
		{
			name:      "pop from empty stack",
			initStack: ArrayStack[int]{},
			testPop: func(t *testing.T, as *ArrayStack[int]) []popResulst {
				result, ok := as.Pop()
				return []popResulst{{value: result, ok: ok}}
			},
			expectedInStack:       []int{},
			expectedPoppedResults: []popResulst{{value: 0, ok: false}},
		},
		{
			name:      "pop until empty",
			initStack: ArrayStack[int]{elements: []int{1, 2}},
			testPop: func(t *testing.T, as *ArrayStack[int]) []popResulst {
				results := make([]popResulst, 0, 2)
				for !as.IsEmpty() {
					result, ok := as.Pop()
					results = append(results, popResulst{value: result, ok: ok})
				}
				return results
			},
			expectedInStack:       []int{},
			expectedPoppedResults: []popResulst{{value: 2, ok: true}, {value: 1, ok: true}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stack := tc.initStack
			poppedResults := tc.testPop(t, &stack)

			if !slices.Equal(tc.expectedInStack, stack.elements) {
				t.Errorf("expected stack elements %v, got %v", tc.expectedInStack, stack.elements)
			}

			if !slices.EqualFunc(tc.expectedPoppedResults, poppedResults, func(a, b popResulst) bool {
				return a.value == b.value && a.ok == b.ok
			}) {
				t.Errorf("expected popped results %v, got %v", tc.expectedPoppedResults, poppedResults)
			}
		})
	}
}

func TestArrayStackPeek(t *testing.T) {
	type peekResult struct {
		value int
		ok    bool
	}

	testCases := []struct {
		name            string
		initStack       ArrayStack[int]
		expectedPeek    peekResult
		expectedInStack []int
	}{
		{
			name:            "peek on empty stack",
			initStack:       ArrayStack[int]{},
			expectedPeek:    peekResult{value: 0, ok: false},
			expectedInStack: []int{},
		},
		{
			name:            "peek on stack with one element",
			initStack:       ArrayStack[int]{elements: []int{1}},
			expectedPeek:    peekResult{value: 1, ok: true},
			expectedInStack: []int{1},
		},
		{
			name:            "peek on stack with multiple elements",
			initStack:       ArrayStack[int]{elements: []int{1, 2, 3}},
			expectedPeek:    peekResult{value: 3, ok: true},
			expectedInStack: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stack := tc.initStack
			result, ok := stack.Peek()

			if !slices.Equal(tc.expectedInStack, stack.elements) {
				t.Errorf("expected stack elements %v, got %v", tc.expectedInStack, stack.elements)
			}

			if result != tc.expectedPeek.value || ok != tc.expectedPeek.ok {
				t.Errorf("expected peek result %v, got %v", tc.expectedPeek, peekResult{value: result, ok: ok})
			}
		})
	}
}

func TestArrayStackIsEmpty(t *testing.T) {
	testCases := []struct {
		name      string
		initStack ArrayStack[int]
		expected  bool
	}{
		{
			name:      "empty stack",
			initStack: ArrayStack[int]{},
			expected:  true,
		},
		{
			name:      "non-empty stack",
			initStack: ArrayStack[int]{elements: []int{1}},
			expected:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stack := tc.initStack
			if stack.IsEmpty() != tc.expected {
				t.Errorf("expected IsEmpty to return %v, got %v", tc.expected, stack.IsEmpty())
			}
		})
	}
}
