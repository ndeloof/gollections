package gollections

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSliceIterator(t *testing.T) {
	it := NewSliceIterator([]T{
		"one", "two", "three",
	})
	assert.Check(t, it.HasNext())
	assert.Check(t, it.Next().Equal("one"))
	assert.Check(t, it.HasNext())
	assert.Check(t, it.Next().Equal("two"))
	assert.Check(t, it.HasNext())
	assert.Check(t, it.Next().Equal("three"))
	assert.Check(t, !it.HasNext())
}
