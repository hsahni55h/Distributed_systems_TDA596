package queue

import (
	"errors"
)

// Queue is a generic FIFO queue implementation with a fixed-size buffer.
type Queue[T any] struct {
	buffer []T
	size   int
}

// New creates a new Queue with the specified size.
func New[T any](size int) *Queue[T] {
	return &Queue[T]{
		buffer: make([]T, 0, size),
		size:   size,
	}
}

// Enqueue adds an element to the end of the queue.
func (q *Queue[T]) Enqueue(item T) error {
	if len(q.buffer) == q.size {
		return errors.New("queue is full")
	}

	q.buffer = append(q.buffer, item)
	return nil
}

// Dequeue removes and returns the front element from the queue.
func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.buffer) == 0 {
		var x T
		return x, errors.New("queue is empty")
	}

	item := q.buffer[0]
	q.buffer = q.buffer[1:]
	return item, nil
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.buffer) == 0
}

// IsFull returns true if the queue is full.
func (q *Queue[T]) IsFull() bool {
	return len(q.buffer) == q.size
}

// Size returns the current number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.buffer)
}
