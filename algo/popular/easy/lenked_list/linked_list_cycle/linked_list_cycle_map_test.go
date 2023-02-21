package main

import "testing"

func Test_hasCycleWithMap(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				head: func() *ListNode {
					node4 := &ListNode{
						Val:  -4,
						Next: nil,
					}
					node3 := &ListNode{
						Val:  0,
						Next: node4,
					}
					node2 := &ListNode{
						Val:  2,
						Next: node3,
					}
					node1 := &ListNode{
						Val:  3,
						Next: node2,
					}

					node4.Next = node2 // Cycle
					return node1
				}(),
			},
			want: true, // Есть цикл
		},
		{
			name: "CASE-2",
			args: args{
				head: func() *ListNode {
					node2 := &ListNode{
						Val: 2,
					}
					node1 := &ListNode{
						Val:  1,
						Next: node2,
					}

					node2.Next = node1 // Cycle
					return node1
				}(),
			},
			want: true,
		},
		{
			name: "CASE-3",
			args: args{
				head: &ListNode{
					Val: 1,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycleWithMap(tt.args.head); got != tt.want {
				t.Errorf("hasCycleWithMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
