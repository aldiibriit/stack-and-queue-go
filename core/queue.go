/*
Package core provides a slice-based Queue data structure
*/
package core

import "fmt"

// Queue : data structure
type Queue struct {
	data []int
}

// NewQueue : returns a new instance of a Queue
func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// IsEmpty : checks whether the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek : returns the next element in the queue
func (q *Queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

// Queue : adds an element onto the queue and returns an pointer to the current queue
func (q *Queue) Add(n int) *Queue {
	q.data = append(q.data, n)
	return q
}

// Dequeue : removes the next element from the queue and returns its value
func (q *Queue) Remove() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}
