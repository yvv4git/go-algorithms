package main

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEndV2(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "CASE-1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				n:    2,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
		},
		{
			name: "CASE-2",
			args: args{
				head: &ListNode{1, nil},
				n:    1,
			},
			want: nil,
		},
		{
			name: "CASE-3",
			args: args{
				head: &ListNode{1, &ListNode{2, nil}},
				n:    1,
			},
			want: &ListNode{1, nil},
		},
		{
			name: "CASE-4",
			args: args{
				head: &ListNode{1, &ListNode{2, nil}},
				n:    2,
			},
			want: &ListNode{2, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEndV2(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEndV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
