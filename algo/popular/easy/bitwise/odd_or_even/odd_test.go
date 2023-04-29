package odd_or_even

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isOdd(n int) bool {
	/*
		Функция позволяет определить - не четное ли число.
	*/
	if n&1 != 0 {
		return true
	}

	return false
}

func Test_isOdd(t *testing.T) {
	assert.True(t, isOdd(1))  // нечетное
	assert.False(t, isOdd(2)) // четное
}
