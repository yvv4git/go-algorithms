package odd_or_even

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isEven(n int) bool {
	/*
		Функция позволяет определить - четное ли число.
	*/
	if n&1 == 0 {
		return true
	}

	return false
}

func Test_isEven(t *testing.T) {
	assert.False(t, isEven(1)) // нечетное
	assert.True(t, isEven(2))  // четное
	assert.False(t, isEven(3)) // нечетное
}
