package lifo

import (
	"errors"

	"github.com/aukbit/cache"
)

var (
	ErrStackIsFull = errors.New("stack is full")
)

// Item generic type of an item in the Stack
type Item interface{}

// Node is a single representation of a data structure in the Stack
type Node struct {
	item Item
	next *Node
}

// Stack is a list of data items where the last item in is the first out..
type Stack struct {
	// first node in the stack
	first *Node
	// last node in the stack
	last *Node
	// number of nodes in the stack
	n int
	// capacity of the stack
	c int
}

func New(capacity int) *Stack {
	return &Stack{
		c: capacity,
	}
}

func (s *Stack) Push(i Item) error {
	n := &Node{
		item: i,
	}
	if s.IsEmpty() {
		s.first = n
		s.last = n
		s.n++
		return nil
	}
	if s.Size() == s.c {
		return ErrStackIsFull
	}
	s.first.next = n
	s.first = n
	s.n++
	return nil
}

func (s *Stack) Pop() Item {
	return s.first.item
}

func (s *Stack) IsEmpty() bool {
	return s.first == nil && s.last == nil
}

func (s *Stack) Size() int {
	return s.n
}

// Iterator returns an iterator to this bag that iterates through the items
// in LIFO order.
func (s *Stack) Iterator() Iterator {
	return newIterator(s.first)
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
