package lifo

import (
	"errors"
)

var (
	ErrStackIsFull = errors.New("stack is full")
)

type Item struct {
	Data interface{}
	next *Item
}

type Stack struct {
	first *Item
	last  *Item
	n     int
	c     int
}

func New(capacity int) *Stack {
	return &Stack{
		c: capacity,
	}
}

func (s *Stack) Push(i *Item) error {
	if s.IsEmpty() {
		s.first = i
		s.last = i
		s.n++
		return nil
	}
	if s.Size() == s.c {
		return ErrStackIsFull
	}
	s.first.next = i
	s.first = i
	s.n++
	return nil
}

func (s *Stack) Pop() *Item {
	return s.first
}

func (s *Stack) IsEmpty() bool {
	return s.first == nil && s.last == nil
}

func (s *Stack) Size() int {
	return s.n
}
