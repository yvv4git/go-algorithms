package main

import (
	"reflect"
	"testing"
)

func Test_middleNodeV1(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Case 1",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			},
			want: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		},
		{
			name: "Case 2",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}},
			},
			want: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNodeV1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNodeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
