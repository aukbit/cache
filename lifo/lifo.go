package lifo

import (
	"errors"
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
