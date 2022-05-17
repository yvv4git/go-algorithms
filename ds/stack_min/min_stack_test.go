package stackmin

import "testing"

func TestMinStack_CommonOperations(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	t.Logf("Min: %d", obj.GetMin()) // -3

	obj.Pop()
	t.Logf("Top: %d", obj.Top()) // 0

	t.Logf("Get min: %d", obj.GetMin()) // -2

}
