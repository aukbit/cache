package bag

import "errors"

var (
	ErrBagIsFull = errors.New("bag is full")
)

// Item is a single representation of a data structure in the Bag
type Item struct {
	Data interface{}
	next *Item
}

// Bag is a list of data items..
type Bag struct {
	// first item in the bag
	first *Item
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
func (b *Bag) Add(i *Item) error {
	if b.IsEmpty() {
		b.first = i
		b.n++
		return nil
	}
	if b.Size() == b.c {
		return ErrBagIsFull
	}
	b.first.next = i
	b.first = i
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
