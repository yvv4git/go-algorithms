package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Не самый эффективный метод.
func IsEvenV1(n int) bool {
	return 0 == n%2
}

func IsEvenV2Xor(n int) bool {
	return n^1 == n+1
}

func IsEvenV3And(n int) bool {
	return !(n&1 == 1)
}

func IsEvenV4Or(n int) bool {
	return n|1 > n
}

func TestIsEven001(t *testing.T) {
	assert.True(t, IsEvenV1(2))
	assert.True(t, IsEvenV2Xor(2))
	assert.True(t, IsEvenV3And(2))
	assert.True(t, IsEvenV4Or(2))

	assert.False(t, IsEvenV1(5))
	assert.False(t, IsEvenV2Xor(5))
	assert.False(t, IsEvenV3And(5))
	assert.False(t, IsEvenV4Or(5))
}
