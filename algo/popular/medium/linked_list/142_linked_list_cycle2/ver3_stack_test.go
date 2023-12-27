package _42_linked_list_cycle2

import (
	"reflect"
	"testing"
)

func Test_detectCycleV3(t *testing.T) {
	type args struct {
		head *ListNode
	}

	// Создаем тестовые данные
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // Создаем цикл

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Cycle at node 2",
			args: args{
				head: node1,
			},
			want: node2,
		},
		{
			name: "No cycle",
			args: args{
				head: &ListNode{Val: 1},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycleV3(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycleV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
