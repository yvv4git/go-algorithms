package _328_break_a_palindrome

import "testing"

func Test_breakPalindromeV1(t *testing.T) {
	type args struct {
		palindrome string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				palindrome: "a",
			},
			want: "",
		},
		{
			name: "Test case 2",
			args: args{
				palindrome: "aa",
			},
			want: "ab",
		},
		{
			name: "Test case 3",
			args: args{
				palindrome: "aba",
			},
			want: "aaa",
		},
		{
			name: "Test case 4",
			args: args{
				palindrome: "abccba",
			},
			want: "aaccba",
		},
		//{
		//	name: "Test case 5",
		//	args: args{
		//		palindrome: "abcba",
		//	},
		//	want: "abdba",
		//},
		//{
		//	name: "Test case 6",
		//	args: args{
		//		palindrome: "abcdba",
		//	},
		//	want: "abceba",
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakPalindromeV1(tt.args.palindrome); got != tt.want {
				t.Errorf("breakPalindromeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
