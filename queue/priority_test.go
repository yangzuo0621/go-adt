package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	// Create a new priority queue
	pq := NewPriorityQueue[int](func(a, b int) bool {
		return a < b
	})

	// Test IsEmpty() on an empty queue
	assert.True(t, pq.IsEmpty(), "Expected IsEmpty() to return true for an empty queue")

	// Test Enqueue() and Dequeue() on a non-empty queue
	pq.Enqueue(5)
	pq.Enqueue(3)
	pq.Enqueue(7)

	// Test IsEmpty() on a non-empty queue
	assert.False(t, pq.IsEmpty(), "Expected IsEmpty() to return false for a non-empty queue")

	// Test Dequeue() on a non-empty queue
	item := pq.Dequeue()
	assert.Equal(t, 3, item, "Expected Dequeue() to return 3, got %d", item)

	// Test Dequeue() on a non-empty queue
	assert.Equal(t, 5, pq.Dequeue(), "Expected Dequeue() to return 5, got %d", item)

	// Test Dequeue() on a non-empty queue
	item = pq.Dequeue()
	assert.Equal(t, 7, item, "Expected Dequeue() to return 7, got %d", item)

	// Test Dequeue() on an empty queue
	assert.True(t, pq.IsEmpty(), "Expected IsEmpty() to return true for an empty queue")
}
