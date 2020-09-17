package gollections

// Iterator define an iterator
type Iterator interface {
	HasNext() bool
	Next() T
}

// AbstractIterator make it easy to define Iterators by only implementing a ComputeNextFn
type AbstractIterator struct {
	hasNext     bool
	computeNext ComputeNextFn
}

// ComputeNextFn compute the next item for an iterator and invoke done when none left
type ComputeNextFn func(done func()) T

// NewSliceIterator build an iterator for a slice
func NewSliceIterator(slice []T) Iterator {
	idx := 0
	next := func(done func()) T {
		v := slice[idx]
		idx = idx + 1
		if idx == len(slice) {
			done()
		}
		return v
	}
	return &AbstractIterator{
		hasNext:     len(slice) > 0,
		computeNext: next,
	}
}

func (it *AbstractIterator) HasNext() bool {
	return it.hasNext
}

func (it *AbstractIterator) Next() T {
	return it.computeNext(it.done)
}

func (it *AbstractIterator) done() {
	it.hasNext = false
}
