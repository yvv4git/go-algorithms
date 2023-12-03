package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoListsV2(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				l1: func() *ListNode {
					result := NewListNode(1)
					result.AddNodeWithValue(2).
						AddNodeWithValue(4)
					return result
				}(),
				l2: func() *ListNode {
					result := NewListNode(1)
					result.AddNodeWithValue(3).
						AddNodeWithValue(4)
					return result
				}(),
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "CASE-2",
			args: args{
				l1: func() *ListNode {
					return nil
				}(),
				l2: func() *ListNode {
					result := NewListNode(0)
					return result
				}(),
			},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoListsV2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoListsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
