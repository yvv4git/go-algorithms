package _2_reverse_linked_list_II

import (
	"reflect"
	"testing"
)

func Test_reverseBetweenV1(t *testing.T) {
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
			if got := reverseBetweenV1(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetweenV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseBetweenV1(t *testing.T) {
	/*
		Альтернативный тест, сгенерированный с помощью ИИ.
	*/
	// Создаем тестовый список: 1->2->3->4->5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	// Вызываем функцию reverseBetweenV1
	left := 2
	right := 4
	result := reverseBetweenV1(head, left, right)

	// Проверяем результат
	expected := []int{1, 4, 3, 2, 5}
	i := 0
	for result != nil {
		if result.Val != expected[i] {
			t.Errorf("Expected value at position %d to be %d, but got %d", i, expected[i], result.Val)
		}
		result = result.Next
		i++
	}
}
