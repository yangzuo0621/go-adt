package queue

// Option is a functional option for configuring a queue.
type Option[T any] func(*queue[T])

// WithCapacity sets the initial capacity of the queue.
func WithCapacity[T any](capacity int) Option[T] {
	return func(q *queue[T]) {
		q.items = make([]T, 0, capacity)
	}
}

// WithInitialItems sets the initial items of the queue.
func WithInitialItems[T any](items []T) Option[T] {
	return func(q *queue[T]) {
		q.items = items
	}
}

func New[T any](options ...Option[T]) *queue[T] {
	q := queue[T]{}
	for _, option := range options {
		option(&q)
	}
	return &q
}

type queue[T any] struct {
	items []T
}

// IsEmpty returns true if the queue is empty.
func (q *queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Enqueue adds an item to the end of the queue.
func (q *queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes the first item from the queue and returns it.
// Note: should use IsEmpty to check if the queue is empty before calling Dequeue.
// Otherwise, it will return the zero value of the type T, which could be misleading for empty queues.
func (q *queue[T]) Dequeue() T {
	var item T
	if q.IsEmpty() {
		return item
	}
	item = q.items[0]
	q.items = q.items[1:]
	return item
}
