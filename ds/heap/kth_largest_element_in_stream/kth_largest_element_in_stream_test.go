package kth_largest_element_in_stream

import "testing"

func TestFindKthElementInStream(t *testing.T) {
	k := 3
	arr := []int{4, 5, 8, 2}
	kElem := Constructor(k, arr)

	v := 3
	t.Logf("v=%d result=%d", v, kElem.Add(v)) // 4

	v = 5
	t.Logf("v=%d result=%d", v, kElem.Add(v)) // 5

	v = 10
	t.Logf("v=%d result=%d", v, kElem.Add(v)) // 5

	v = 9
	t.Logf("v=%d result=%d", v, kElem.Add(v)) // 8

	v = 4
	t.Logf("v=%d result=%d", v, kElem.Add(v)) // 8
}
