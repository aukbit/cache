package fifo

import (
	"errors"
)

var (
	ErrStackIsFull = errors.New("stack is full")
)

// Item generic type of an item in this bag
type Item interface{}

// Node is a single representation of a data structure in the Queue
type Node struct {
	item Item
	next *Node
}

// Queue is a list of data items where the first item in is first out..
type Queue struct {
	// first node in the queue
	first *Node
	// last node in the queue
	last *Node
	// number of nodes in the queue
	n int
	// capacity of the queue
	c int
}

// New instanciates a new Queue
func New(capacity int) *Queue {
	return &Queue{
		c: capacity,
	}
}

// Enqueue add an item of data to the queue.
func (q *Queue) Enqueue(i Item) error {
	n := &Node{
		item: i,
	}
	if q.IsEmpty() {
		q.first = n
		q.last = n
		q.n++
		return nil
	}
	if q.Size() == q.c {
		return ErrStackIsFull
	}
	q.first.next = n
	q.first = n
	q.n++
	return nil
}

// Dequeue remove and returns the least recently added item from the queue.
func (q *Queue) Dequeue() Item {
	if q.IsEmpty() {
		return nil
	}
	last := q.last
	// last node
	if q.last.next == nil {
		q.first = nil
	}
	q.last = q.last.next
	q.n--
	return last.item
}

// IsEmpty is the queue empty?
func (q *Queue) IsEmpty() bool {
	return q.first == nil && q.last == nil
}

// Size number of items in the queue.
func (q *Queue) Size() int {
	return q.n
}

// Capacity number of items possible in the queue.
func (q *Queue) Capacity() int {
	return q.c
}
