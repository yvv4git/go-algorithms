package main

import "testing"

func Test_reorderListV2(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test Case 1: Normal list",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "Test Case 2: Single node list",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
		{
			name: "Test Case 3: Empty list",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderListV2(tt.args.head)
		})
	}
}
