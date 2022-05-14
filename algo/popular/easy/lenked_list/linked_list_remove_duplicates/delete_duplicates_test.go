package main

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				head: func() *ListNode {
					node3 := ListNode{
						Val: 2,
					}
					node2 := ListNode{
						Val:  1,
						Next: &node3,
					}
					node1 := ListNode{
						Val:  1,
						Next: &node2,
					}
					return &node1
				}(),
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-2",
			args: args{
				head: func() *ListNode {
					node3 := ListNode{
						Val: 2,
					}
					node2 := ListNode{
						Val:  1,
						Next: &node3,
					}
					node1 := ListNode{
						Val:  1,
						Next: &node2,
					}
					return &node1
				}(),
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-3",
			args: args{
				head: func() *ListNode {
					node5 := ListNode{
						Val: 3,
					}
					node4 := ListNode{
						Val:  3,
						Next: &node5,
					}
					node3 := ListNode{
						Val:  2,
						Next: &node4,
					}
					node2 := ListNode{
						Val:  1,
						Next: &node3,
					}
					node1 := ListNode{
						Val:  1,
						Next: &node2,
					}
					return &node1
				}(),
			},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := deleteDuplicates(tt.args.head)
			resultInt := []int{}
			for {
				resultInt = append(resultInt, result.Val)

				if result.Next == nil {
					break
				}
				result = result.Next
			}

			if !reflect.DeepEqual(resultInt, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", resultInt, tt.want)
			}
		})
	}
}
