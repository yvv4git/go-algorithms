package _4_swap_nodes_in_pairs

import (
	"reflect"
	"testing"
)

func Test_swapPairsV1(t *testing.T) {
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
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			},
			want: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
		},
		{
			name: "Test Case 2",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			},
			want: &ListNode{2, &ListNode{1, &ListNode{3, nil}}},
		},
		{
			name: "Test Case 3",
			args: args{
				head: &ListNode{1, nil},
			},
			want: &ListNode{1, nil},
		},
		{
			name: "Test Case 4",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairsV1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
