package _05_design_hashset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVer2(t *testing.T) {
	obj := ConstructorV2()
	obj.Add(1)
	obj.Add(2)

	assert.True(t, obj.Contains(1))
	assert.False(t, obj.Contains(3))

	obj.Add(2)
	assert.True(t, obj.Contains(2))

	obj.Remove(2)
	assert.False(t, obj.Contains(2))
}
