package queue

func NewPriorityQueue[T any](compare func(a, b T) bool) *priorityQueue[T] {
	return &priorityQueue[T]{
		compare: compare,
	}
}

type priorityQueue[T any] struct {
	items   []T
	compare func(a, b T) bool
}

// IsEmpty returns true if the queue is empty.
func (q *priorityQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Enqueue adds an item to the queue.
func (q *priorityQueue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
	j := len(q.items) - 1
	for j > 0 {
		i := (j - 1) / 2
		if q.compare(q.items[j], q.items[i]) {
			q.items[i], q.items[j] = q.items[j], q.items[i]
			j = i
		} else {
			break
		}
	}
}

// Dequeue removes the first item from the queue and returns it.
func (q *priorityQueue[T]) Dequeue() T {
	if q.IsEmpty() {
		var item T
		return item
	}

	item := q.items[0]
	q.items[0] = q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	q.heapify(0)

	return item
}

// heapify maintains the heap property of the queue.
func (q *priorityQueue[T]) heapify(p int) {
	left := 2*p + 1
	right := 2*p + 2
	target := p
	if left < len(q.items) && q.compare(q.items[left], q.items[p]) {
		target = left
	}
	if right < len(q.items) && q.compare(q.items[right], q.items[target]) {
		target = right
	}
	if target != p {
		q.items[p], q.items[target] = q.items[target], q.items[p]
		q.heapify(target)
	}
}
