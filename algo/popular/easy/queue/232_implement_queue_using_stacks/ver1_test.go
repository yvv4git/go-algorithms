package _32_implement_queue_using_stacks

import "testing"

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	param2 := obj.Pop()
	param3 := obj.Peek()
	param4 := obj.Empty()
	t.Logf("Param-2: %v", param2)
	t.Logf("Param-3: %v", param3)
	t.Logf("Param-4: %v", param4)
}
