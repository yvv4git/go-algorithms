package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
			desc: "We have palindrome: 'amanaplanacanalpanama'",
		},
		{
			name: "CASE-2",
			args: args{
				s: "race a car",
			},
			want: false,
			desc: "'raceacar' is not palindrome",
		},
		{
			name: "CASE-3",
			args: args{
				s: " ",
			},
			want: true,
			desc: "s is an empty string '' after removing non-alphanumeric characters. Since an empty string reads the same forward and backward, it is a palindrome.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
