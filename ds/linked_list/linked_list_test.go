package linked_list

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	type args struct {
		datas []string
	}

	tests := []struct {
		name  string
		args  args
		check func(list *LinkedList) bool
	}{
		{
			name: "CASE-1",
			args: args{
				datas: []string{"1", "2", "3"},
			},
			check: func(list *LinkedList) bool {
				firstElement := list.Head
				if firstElement.Data != "1" {
					return false
				}

				secondElement := firstElement.Next
				if secondElement.Data != "2" {
					return false
				}

				return list.Tail.Data == "3"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewLinkedList(tt.args.datas)
			if !tt.check(result) {
				t.Error("NewLinkedList() works incorrect")
			}
		})
	}
}

func TestLinkedList_Add(t *testing.T) {
	type args struct {
		data string
	}

	tests := []struct {
		name     string
		list     *LinkedList
		args     args
		wantTail string
	}{
		{
			name: "CASE-1",
			list: NewLinkedList([]string{"1", "2", "3"}),
			args: args{
				data: "4",
			},
			wantTail: "4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Add(tt.args.data)
			if tt.list.Tail.Data != tt.wantTail {
				t.Errorf("The last element should be %s, but there %s", tt.wantTail, tt.list.Tail.Data)
			}
		})
	}
}

func TestLinkedList_Delete(t *testing.T) {
	tests := []struct {
		name    string
		prepare func() (*LinkedList, *Element)
		check   func(list *LinkedList) bool
		desc    string
	}{
		{
			name: "CASE-1",
			prepare: func() (*LinkedList, *Element) {
				list := NewLinkedList([]string{"1", "2", "3"})
				return list, list.Head
			},
			check: func(list *LinkedList) bool {
				firstElement := list.Head
				return firstElement.Data == "2"
			},
			desc: "Delete first element from list",
		},
		{
			name: "CASE-2",
			prepare: func() (*LinkedList, *Element) {
				list := NewLinkedList([]string{"1", "2", "3"})
				return list, list.Tail
			},
			check: func(list *LinkedList) bool {
				lastElement := list.Tail
				return lastElement.Data == "2"
			},
			desc: "Delete last element from list",
		},
		{
			name: "CASE-3",
			prepare: func() (*LinkedList, *Element) {
				list := NewLinkedList([]string{"1", "2", "3"})
				return list, list.Head.Next
			},
			check: func(list *LinkedList) bool {
				return list.Head.Next == list.Tail
			},
			desc: "Delete middle element",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, element := tt.prepare()
			list.Delete(element)
			if !tt.check(list) {
				t.Error("We have problems with check Delete()")
			}
		})
	}
}

func TestLinkedList_InsertAfter(t *testing.T) {
	tests := []struct {
		name    string
		prepare func() (list *LinkedList, insert, after *Element)
		check   func(list *LinkedList) bool
		desc    string
	}{
		{
			name: "CASE-1",
			prepare: func() (list *LinkedList, insert, after *Element) {
				list = NewLinkedList([]string{"1", "2", "3"})
				insertElement := &Element{
					Data: "A",
				}
				return list, insertElement, list.Head
			},
			check: func(list *LinkedList) bool {
				elementAfterInsert := list.Head.Next.Next
				return elementAfterInsert.Data == "2"
			},
			desc: "When you add after the first element",
		},
		{
			name: "CASE-2",
			prepare: func() (list *LinkedList, insert, after *Element) {
				list = NewLinkedList([]string{"1", "2", "3"})
				insertElement := &Element{
					Data: "A",
				}
				return list, insertElement, list.Head.Next
			},
			check: func(list *LinkedList) bool {
				elementAfterInsert := list.Head.Next.Next
				return elementAfterInsert.Data == "A"
			},
			desc: "When you add element after middle element",
		},
		{
			name: "CASE-3",
			prepare: func() (list *LinkedList, insert, after *Element) {
				list = NewLinkedList([]string{"1", "2", "3"})
				insertElement := &Element{
					Data: "A",
				}
				return list, insertElement, list.Tail
			},
			check: func(list *LinkedList) bool {
				elementAfterInsert := list.Tail
				return elementAfterInsert.Data == "A"
			},
			desc: "When you add element in the end of list",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, insert, after := tt.prepare()
			list.InsertAfter(insert, after)
			if !tt.check(list) {
				t.Error("We have problems with check InsertAfter()")
			}
		})
	}
}
