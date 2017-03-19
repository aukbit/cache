package cache

import "errors"

var (
	ErrNoSuchElement        = errors.New("no more elements in the iterator")
	ErrUnsupportedOperation = errors.New("unsupported operation")
)
