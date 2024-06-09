package _2_reverse_linked_list_II

import (
	"reflect"
	"testing"
)

func Test_reverseBetweenV4(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "CASE-1",
			args: args{
				head: func() *ListNode {
					// Создаем тестовый список: 1->2->3->4->5
					head := &ListNode{Val: 1}
					head.Next = &ListNode{Val: 2}
					head.Next.Next = &ListNode{Val: 3}
					head.Next.Next.Next = &ListNode{Val: 4}
					head.Next.Next.Next.Next = &ListNode{Val: 5}
					return head
				}(),
				left:  2,
				right: 4,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "CASE-2",
			args: args{
				head: func() *ListNode {
					// Создаем тестовый список: 5
					head := &ListNode{Val: 5}
					return head
				}(),
				left:  1,
				right: 1,
			},
			want: &ListNode{
				Val: 5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetweenV4(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetweenV4() = %v, want %v", got, tt.want)
			}
		})
	}
}
