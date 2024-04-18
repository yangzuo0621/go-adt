package stack

// Option is a functional option for configuring a stack.
type Option[T any] func(*stack[T])

// WithCapacity sets the initial capacity of the stack.
func WithCapacity[T any](capacity int) Option[T] {
	return func(s *stack[T]) {
		s.items = make([]T, 0, capacity)
	}
}

func New[T any](options ...Option[T]) *stack[T] {
	s := stack[T]{}
	for _, option := range options {
		option(&s)
	}
	return &s
}

type stack[T any] struct {
	items []T
}

// Push adds an item to the top of the stack.
func (s *stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Push adds an item to the top of the stack.
func (s *stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes the top item from the stack and returns it.
// Note: should use IsEmpty to check if the stack is empty before calling Pop.
// Otherwise, it will return the zero value of the type T, which could be misleading for empty stacks.
func (s *stack[T]) Pop() T {
	var item T
	if s.IsEmpty() {
		return item
	}
	item = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the top item from the stack without removing it.
func (s *stack[T]) Peek() T {
	var item T
	if s.IsEmpty() {
		return item
	}
	return s.items[len(s.items)-1]
}

// Size returns the number of items in the stack.
func (s *stack[T]) Size() int {
	return len(s.items)
}
