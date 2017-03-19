package fifo

import (
	"errors"

	"github.com/aukbit/cache"
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

// Iterator returns an iterator to this bag that iterates through the items in arbitrary order.
func (q *Queue) Iterator() Iterator {
	return newIterator(q.first)
}

// Iterator represents an iterator over a collection.
type Iterator struct {
	current *Node
}

func newIterator(n *Node) Iterator {
	return Iterator{
		current: n,
	}
}

// HasNext returns true if the iteration has more elements.
func (i *Iterator) HasNext() bool {
	return i.current != nil
}

// Remove removes from the underlying collection the last element returned by the iterator (optional operation).
func (i *Iterator) Remove() error {
	return cache.ErrUnsupportedOperation
}

// Next returns the next element in the iteration.
func (i *Iterator) Next() (Item, error) {
	if !i.HasNext() {
		return nil, cache.ErrNoSuchElement
	}
	item := i.current.item
	i.current = i.current.next
	return item, nil
}
