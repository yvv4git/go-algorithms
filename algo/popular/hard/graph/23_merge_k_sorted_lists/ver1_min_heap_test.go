package _3_merge_k_sorted_lists

import (
	"reflect"
	"testing"
)

func Test_mergeKListsV1(t *testing.T) {
	type args struct {
		lists []*ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test Case 1",
			args: args{
				lists: []*ListNode{
					newListNode(1, 4, 5),
					newListNode(1, 3, 4),
					newListNode(2, 6),
				},
			},
			want: newListNode(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			name: "Test Case 2",
			args: args{
				lists: []*ListNode{
					newListNode(1, 2, 3),
					newListNode(4, 5, 6),
					newListNode(7, 8, 9),
				},
			},
			want: newListNode(1, 2, 3, 4, 5, 6, 7, 8, 9),
		},
		{
			name: "Test Case 3",
			args: args{
				lists: []*ListNode{
					newListNode(1),
					newListNode(2),
					newListNode(3),
				},
			},
			want: newListNode(1, 2, 3),
		},
		{
			name: "Test Case 4",
			args: args{
				lists: []*ListNode{
					newListNode(),
					newListNode(),
					newListNode(),
				},
			},
			want: newListNode(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKListsV1(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKListsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newListNode(values ...int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	return head
}
