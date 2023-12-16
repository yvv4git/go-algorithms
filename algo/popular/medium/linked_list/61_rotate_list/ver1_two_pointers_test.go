package _1_rotate_list

import (
	"reflect"
	"testing"
)

func Test_rotateRightV1(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test Case 1",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
				k:    2,
			},
			want: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{
			name: "Test Case 2",
			args: args{
				head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
				k:    4,
			},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			name: "Test Case 3",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
				k:    0,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			name: "Test Case 4",
			args: args{
				head: &ListNode{Val: 1},
				k:    1,
			},
			want: &ListNode{Val: 1},
		},
		{
			name: "Test Case 5",
			args: args{
				head: nil,
				k:    1,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRightV1(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRightV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
