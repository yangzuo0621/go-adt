package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("WithCapacity", func(t *testing.T) {
		q := New[int](WithCapacity[int](3))
		assert.Equal(t, 0, len(q.items), "expected length to be 0")
		assert.Equal(t, 3, cap(q.items), "expected capacity to be 3")
	})

	t.Run("WithInitialItems", func(t *testing.T) {
		q := New[int](WithInitialItems[int]([]int{1, 2, 3}))
		assert.Equal(t, []int{1, 2, 3}, q.items, "expected items to be [1, 2, 3]")
	})

	t.Run("IsEmpty", func(t *testing.T) {
		q := New[int]()
		assert.True(t, q.IsEmpty(), "expected queue to be empty")

		q = New[int](WithInitialItems[int]([]int{1, 2, 3}))
		assert.False(t, q.IsEmpty(), "expected queue not to be empty")

		_ = q.Dequeue()
		_ = q.Dequeue()
		_ = q.Dequeue()
		assert.True(t, q.IsEmpty(), "expected queue to be empty")
	})

	t.Run("Enqueue", func(t *testing.T) {
		q := New[int]()
		q.Enqueue(1)
		assert.False(t, q.IsEmpty(), "expected queue not to be empty")
		assert.Equal(t, 1, q.items[0], "expected item to be 1")
		assert.Equal(t, 1, len(q.items), "expected length to be 1")
	})

	t.Run("Dequeue", func(t *testing.T) {
		q := New[int](WithInitialItems[int]([]int{1, 2, 3}))
		assert.Equal(t, 1, q.Dequeue(), "expected item to be 1")
		assert.Equal(t, 2, q.Dequeue(), "expected item to be 2")
		assert.Equal(t, 3, q.Dequeue(), "expected item to be 3")
		assert.Equal(t, 0, len(q.items), "expected length to be 0")
	})

	t.Run("Peek", func(t *testing.T) {
		q := New[int](WithInitialItems[int]([]int{1, 2, 3}))
		assert.Equal(t, 1, q.Peek(), "expected item to be 1")
		assert.Equal(t, 1, q.Dequeue(), "expected item to be 1")
		assert.Equal(t, 2, q.Peek(), "expected item to be 2")
		assert.Equal(t, 2, q.Dequeue(), "expected item to be 2")
		assert.Equal(t, 3, q.Peek(), "expected item to be 3")
		assert.Equal(t, 3, q.Dequeue(), "expected item to be 3")
		assert.Equal(t, 0, len(q.items), "expected length to be 0")
	})

	t.Run("Len", func(t *testing.T) {
		q := New[int]()
		assert.Equal(t, 0, q.Len(), "expected length to be 0")

		q.Enqueue(1)
		assert.Equal(t, 1, q.Len(), "expected length to be 1")

		_ = q.Dequeue()
		assert.Equal(t, 0, q.Len(), "expected length to be 0")
	})
}
