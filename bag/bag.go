package bag

import (
	"errors"

	"github.com/aukbit/cache"
)

var (
	ErrBagIsFull = errors.New("bag is full")
)

// Item generic type of an item in this bag
type Item interface{}

// Node is a single representation of a data structure in the Bag
type Node struct {
	item Item
	next *Node
}

// Bag is a list of data items..
type Bag struct {
	// first item in the bag
	first *Node
	// number of items in the bag
	n int
	// capacity of the bag
	c int
}

// New create a new Bag
func New(capacity int) *Bag {
	return &Bag{
		c: capacity,
	}
}

// Add an item
func (b *Bag) Add(i Item) error {
	n := &Node{
		item: i,
	}
	if b.IsEmpty() {
		b.first = n
		b.n++
		return nil
	}
	if b.Size() == b.c {
		return ErrBagIsFull
	}
	b.first.next = n
	b.first = n
	b.n++
	return nil
}

// IsEmpty is the bag empty?
func (b *Bag) IsEmpty() bool {
	return b.first == nil
}

// Size number of items in the bag
func (b *Bag) Size() int {
	return b.n
}

// Iterator returns an iterator to this bag that iterates through the items in arbitrary order.
func (b *Bag) Iterator() Iterator {
	return newIterator(b.first)
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
