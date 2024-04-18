package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("WithCapacity", func(t *testing.T) {
		s := New[int](WithCapacity[int](3))
		assert.Equal(t, 0, len(s.items), "expected length to be 0")
		assert.Equal(t, 3, cap(s.items), "expected capacity to be 3")
	})

	t.Run("IsEmpty", func(t *testing.T) {
		s := New[int]()
		assert.True(t, s.IsEmpty(), "expected stack to be empty")

		s = New[int](WithCapacity[int](3))
		assert.True(t, s.IsEmpty(), "expected stack to be empty")
	})

	t.Run("Push", func(t *testing.T) {
		s := New[int]()
		s.Push(1)
		assert.False(t, s.IsEmpty(), "expected stack not to be empty")
		assert.Equal(t, 1, s.items[0], "expected item to be 1")
		assert.Equal(t, 1, len(s.items), "expected length to be 1")
	})

	t.Run("Pop", func(t *testing.T) {
		s := New[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		assert.Equal(t, 3, s.Pop(), "expected item to be 3")
		assert.Equal(t, 2, s.Pop(), "expected item to be 2")
		assert.Equal(t, 1, s.Pop(), "expected item to be 1")
		assert.Equal(t, 0, len(s.items), "expected length to be 0")
	})

	t.Run("Peek", func(t *testing.T) {
		s := New[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		assert.Equal(t, 3, s.Peek(), "expected item to be 3")
		assert.Equal(t, 3, s.Pop(), "expected item to be 3")
		assert.Equal(t, 2, s.Peek(), "expected item to be 2")
		assert.Equal(t, 2, s.Pop(), "expected item to be 2")
		assert.Equal(t, 1, s.Peek(), "expected item to be 1")
		assert.Equal(t, 1, s.Pop(), "expected item to be 1")
		assert.Equal(t, 0, len(s.items), "expected length to be 0")
	})

	t.Run("Len", func(t *testing.T) {
		s := New[int]()
		assert.Equal(t, 0, s.Len(), "expected length to be 0")

		s.Push(1)
		assert.Equal(t, 1, s.Len(), "expected length to be 1")

		s.Push(2)
		assert.Equal(t, 2, s.Len(), "expected length to be 2")

		_ = s.Pop()
		assert.Equal(t, 1, s.Len(), "expected length to be 1")

		_ = s.Pop()
		assert.Equal(t, 0, s.Len(), "expected length to be 0")
	})
}
