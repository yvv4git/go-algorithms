package heap_int_max

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()

	assert.Equal(t, 0, h.Len())

	h.Push(46)
	h.Push(86)
	h.Push(6)
	h.Push(66)
	h.Push(96)
	h.Push(16)

	assert.Equal(t, 6, h.Len())

	assert.Equal(t, 96, h.Peek())
	assert.Equal(t, 96, h.Pop())
	assert.Equal(t, 86, h.Peek())
	assert.Equal(t, 86, h.Pop())
	assert.Equal(t, 66, h.Peek())
	assert.Equal(t, 66, h.Pop())
	assert.Equal(t, 46, h.Peek())
	assert.Equal(t, 46, h.Pop())
	assert.Equal(t, 16, h.Peek())
	assert.Equal(t, 16, h.Pop())
	assert.Equal(t, 6, h.Peek())
	assert.Equal(t, 6, h.Pop())

	assert.Equal(t, 0, h.Len())
}
