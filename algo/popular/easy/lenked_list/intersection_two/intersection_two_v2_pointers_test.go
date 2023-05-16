package intersection_two

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNodeV2(t *testing.T) {
	type testOptions struct {
		headA *ListNode
		headB *ListNode
		want  *ListNode
	}

	tests := []struct {
		name    string
		options testOptions
	}{
		{
			name: "CASE-1",
			options: func() testOptions {
				common := &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				}

				headA := &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  1,
						Next: common,
					},
				}

				headB := &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  1,
							Next: common,
						},
					},
				}

				return testOptions{
					headA: headA,
					headB: headB,
					want:  common,
				}
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNodeV2(tt.options.headA, tt.options.headB); !reflect.DeepEqual(got, tt.options.want) {
				t.Errorf("getIntersectionNodeV2() = %v, want %v", got, tt.options)
			}
		})
	}
}
