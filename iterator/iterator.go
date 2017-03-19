package iterator

import "errors"

var (
	ErrNoSuchElement = errors.New("no more elements in the iterator")
)

// Iterator An iterator over a collection.
// Iterators allow the caller to remove elements from the underlying
// collection during the iteration with well-defined semantics.
type Iterator interface {
	// HasNext returns true if the iteration has more elements.
	HasNext() bool
	// Next returns the next element in the iteration.
	Next() interface{}
	// Remove removes from the underlying collection the last element returned
	// by the iterator (optional operation).
	Remove()
}
