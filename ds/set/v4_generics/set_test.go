package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet[string]()
	s.Add("Vladimir")

	assert.True(t, s.Contains("Vladimir"))
}

func TestSeMultipleValues(t *testing.T) {
	s := NewSet[int](1, 2, 3)
	s.Add(4)
	s.Add(5)

	assert.True(t, s.Contains(4))
}

func TestSet_Members(t *testing.T) {
	s := NewSet[int](1, 2, 3)
	s.Add(4)
	s.Add(5)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, s.Members())
}

func TestSet_String(t *testing.T) {
	s := NewSet[int](1, 2, 3)
	s.Add(4)
	s.Add(5)

	// Так как данные в мапе хранятся в случайном порядке, то в строке они тоже будут в случайном порядке
	t.Logf("=>%s", s.String())
}

func TestSet_Union(t *testing.T) {
	s := NewSet[int](1, 2, 3)
	s.Add(4)
	s.Add(5)

	s2 := NewSet[int](4, 5, 6)
	s2.Add(7)
	s2.Add(8)

	s3 := s.Union(s2)

	t.Logf("Result: %#v", s3)
}

func TestSet_Intersection(t *testing.T) {
	s := NewSet[int](1, 2, 3)
	s2 := NewSet[int](3, 4, 5)
	s3 := s.Intersection(s2)

	t.Logf("s3(intersection) = %#v", s3)
}
