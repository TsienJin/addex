package utils

import "errors"

// Stack is a generic stack that holds type T
type Stack[T any] struct {
	items []T
}

// Push Method to push items to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop take the first element of the stack
func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) <= 0 {
		var emptyVar T
		return emptyVar, errors.New("stack is empty")
	}

	indexOfItem := len(s.items) - 1
	item := s.items[indexOfItem]
	s.items = s.items[:indexOfItem]

	return item, nil
}

// Peek method to see the next item
func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) <= 0 {
		var emptyVar T
		return emptyVar, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty method to check if the stack is empty or not
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) > 0
}
