package fifo

import "fmt"

// Item is a single representation of a data structure in the Queue
type Item struct {
	Data interface{}
	next *Item
}

// Queue is a list of data items.
type Queue struct {
	// first item in the queue
	first *Item
	// last item in the queue
	last *Item
	// number of items in the queue
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
func (q *Queue) Enqueue(i *Item) error {
	if q.IsEmpty() {
		q.first = i
		q.last = i
		q.n++
		return nil
	}
	if q.Size() == q.c {
		return fmt.Errorf("queue is full")
	}
	q.first.next = i
	q.first = i
	q.n++
	return nil
}

// Dequeue remove and returns the least recently added item from the queue.
func (q *Queue) Dequeue() *Item {
	if q.IsEmpty() {
		return nil
	}
	last := q.last
	// last item
	if q.last.next == nil {
		q.first = nil
	}
	q.last = q.last.next
	q.n--
	return last
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
