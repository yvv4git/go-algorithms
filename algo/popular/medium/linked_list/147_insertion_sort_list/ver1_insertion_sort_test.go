package main

import (
	"reflect"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Empty list",
			args: args{head: nil},
			want: nil,
		},
		{
			name: "Single element list",
			args: args{head: &ListNode{Val: 1}},
			want: &ListNode{Val: 1},
		},
		{
			name: "Sorted list",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name: "Unsorted list",
			args: args{head: &ListNode{Val: 3, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name: "Duplicate elements list",
			args: args{head: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
