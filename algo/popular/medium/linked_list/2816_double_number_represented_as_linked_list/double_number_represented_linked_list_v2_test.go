package _816_double_number_represented_as_linked_list

import (
	"reflect"
	"testing"
)

func Test_doubleItV2(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test case 1",
			args: args{
				head: createLinkedList([]int{1, 8, 9}),
			},
			want: createLinkedList([]int{3, 7, 8}),
		},
		{
			name: "Test case 2",
			args: args{
				head: createLinkedList([]int{1, 8, 9}),
			},
			want: createLinkedList([]int{3, 7, 8}),
		},
		{
			name: "Test case 3",
			args: args{
				head: createLinkedList([]int{0, 1, 2, 3, 4}),
			},
			want: createLinkedList([]int{0, 2, 4, 6, 8}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doubleItV2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doubleItV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createLinkedList(values []int) *ListNode {
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
