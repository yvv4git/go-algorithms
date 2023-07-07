package remove_ilnked_list_elements

import (
	"reflect"
	"testing"
)

func Test_removeElementsV2(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "CASE-1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val: 6,
										},
									},
								},
							},
						},
					},
				},
				val: 6,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "CASE-2",
			args: args{
				head: nil,
				val:  1,
			},
			want: nil,
		},
		{
			name: "CASE-3",
			args: args{
				head: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val: 7,
							},
						},
					},
				},
				val: 7,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElementsV2(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElementsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
