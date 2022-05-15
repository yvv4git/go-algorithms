package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				func() *ListNode {
					return &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 1,
								},
							},
						},
					}
				}(),
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				func() *ListNode {
					return &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
						},
					}
				}(),
			},
			want: false,
		},
		{
			name: "CASE-3",
			args: args{
				func() *ListNode {
					return &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
						},
					}
				}(),
			},
			want: true,
		},
		{
			name: "CASE-4",
			args: args{
				func() *ListNode {
					return &ListNode{
						Val: 1,
					}
				}(),
			},
			want: true,
		},
		{
			name: "CASE-5",
			args: args{
				func() *ListNode {
					return &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 1,
							},
						},
					}
				}(),
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecificXOR01(t *testing.T) {
	// s := []int{1, 2, 2, 1}
	s := []int{1, 3, 0, 2}
	x := 0

	t.Logf("x = %d", x)
	for _, val := range s {
		x ^= val
		t.Logf("x[%d] ^= %d", x, val)
	}
}

func TestSpecificXOR02(t *testing.T) {
	x := 1
	t.Logf("x=%d x^%d=%d", x, 0, x^0)
}

// Хочу визуализировать двустороннию итерацию по списку.
func TestVisualyseBeedirectionIterate(t *testing.T) {
	s := []int{1, 2, 3, 2, 1}
	i := 0          // с этого индекса начнем
	j := len(s) - 1 // с этого индекса начнем

	for i < j {
		t.Logf("i=%d j=%d, i==j=>%v", s[i], s[j], s[i] == s[j])
		i++
		j--
	}
	t.Log("stoped")
}
