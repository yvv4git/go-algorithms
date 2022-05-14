package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
			result := mergeTwoLists(tt.args.l1, tt.args.l2)

			resultList := []int{}
			noda := result
			for {
				resultList = append(resultList, noda.Val)
				if noda.Next == nil {
					break
				}

				noda = noda.Next
			}
			if !reflect.DeepEqual(resultList, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", resultList, tt.want)
			}
		})
	}
}
