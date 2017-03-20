package bag

import (
	"github.com/aukbit/cache"
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
}

// New create a new Bag
func New() *Bag {
	return &Bag{}
}

// Add an item
func (b *Bag) Add(i Item) error {
	n := &Node{
		item: i,
		next: b.first,
	}
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
func (b *Bag) Iterator() *Iterator {
	return newIterator(b.first)
}

// Iterator represents an iterator over a collection.
type Iterator struct {
	current *Node
}

func newIterator(n *Node) *Iterator {
	return &Iterator{
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
	// fmt.Println(i.current.item)
	item := i.current.item
	i.current = i.current.next
	// fmt.Println(i.current.next)
	return item, nil
}
