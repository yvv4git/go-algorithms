package _2_remove_duplicates_from_sorted_list2

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicatesV3(t *testing.T) {
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
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val:  4,
									Next: &ListNode{Val: 4},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "Test Case 3",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicatesV3(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicatesV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
