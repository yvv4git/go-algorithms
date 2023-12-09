package _807_insert_greatest_common_divisors_in_linked_list

import (
	"testing"
)

func Test_insertGreatestCommonDivisorsV3(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test Case 1",
			args: args{
				head: &ListNode{
					Val: 18,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 10,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 18,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 10,
								Next: &ListNode{
									Val: 1,
									Next: &ListNode{
										Val: 3,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				head: &ListNode{
					Val: 7,
				},
			},
			want: &ListNode{
				Val: 7,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertGreatestCommonDivisorsV3(tt.args.head); !compareListNodes(got, tt.want) {
				t.Errorf("insertGreatestCommonDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareListNodes(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}
