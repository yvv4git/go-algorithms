package _07_design_linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	fmt.Println(obj.Get(1)) // 2
	obj.DeleteAtIndex(1)
	fmt.Println(obj.Get(1)) // 3
}
