package main

import "testing"

func Test_isPalindromeV4(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1: Palindrome",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Test Case 2: Not Palindrome",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "Test Case 3: Single Node",
			args: args{
				head: &ListNode{
					Val: 1,
				},
			},
			want: true,
		},
		{
			name: "Test Case 4: Empty List",
			args: args{
				head: nil,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeV4(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeV4() = %v, want %v", got, tt.want)
			}
		})
	}
}
