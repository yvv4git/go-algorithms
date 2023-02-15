package queue_on_node

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Constructor()

	for i := 0; i < 5; i++ {
		queue.Push(i)
	}

	var result []int
	for i := 0; i < 5; i++ {
		value := queue.Pop()
		result = append(result, value)
	}

	assert.Equal(t, []int{0, 1, 2, 3, 4}, result)
}
