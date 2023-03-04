package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, queue.items)

	// Берем элемент из начала списка.
	assert.Equal(t, 1, queue.Dequeue())
	assert.Equal(t, 2, queue.Dequeue())
	assert.Equal(t, 3, queue.Dequeue())
	assert.Equal(t, 4, queue.Dequeue())
	assert.Equal(t, 5, queue.Dequeue())
}
