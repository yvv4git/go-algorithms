package queue_on_slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueSimple(t *testing.T) {
	queue := New()
	queue.Push(1)        // Enqueue
	value := queue.Pop() // Dequeue
	assert.Equal(t, 1, value)
}

func TestQueue(t *testing.T) {
	queue := New()
	for i := 0; i < 5; i++ {
		queue.Push(i)
	}

	var result []int
	for i := 0; i < 5; i++ {
		value := queue.Pop()
		result = append(result, value)
	}

	// Проверим, что данные вытаскивали в прямом порядке, как клали(FIFO).
	assert.Equal(t, []int{0, 1, 2, 3, 4}, result)
}
