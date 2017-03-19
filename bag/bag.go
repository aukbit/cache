package bag

import "errors"

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

// func (b *Bag) Iterator() Iterator {
//
// }
