package _66_prime_palindrome

import "testing"

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{n: 2},
			want: true,
		},
		{
			name: "Test 2",
			args: args{n: 3},
			want: true,
		},
		{
			name: "Test 3",
			args: args{n: 4},
			want: false,
		},
		{
			name: "Test 4",
			args: args{n: 5},
			want: true,
		},
		{
			name: "Test 5",
			args: args{n: 10},
			want: false,
		},
		{
			name: "Test 6",
			args: args{n: 13},
			want: true,
		},
		{
			name: "Test 7",
			args: args{n: 15},
			want: false,
		},
		{
			name: "Test 8",
			args: args{n: 21},
			want: false,
		},
		{
			name: "Test 9",
			args: args{n: 23},
			want: true,
		},
		{
			name: "Test 10",
			args: args{n: 25},
			want: false,
		},
		{
			name: "Test 11",
			args: args{n: 10301},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1: Palindrome number",
			args: args{n: 12321},
			want: true,
		},
		{
			name: "Test case 2: Non-palindrome number",
			args: args{n: 12345},
			want: false,
		},
		{
			name: "Test case 3: Single digit number",
			args: args{n: 5},
			want: true,
		},
		{
			name: "Test case 4: Two same digits number",
			args: args{n: 11},
			want: true,
		},
		{
			name: "Test case 5: Zero",
			args: args{n: 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.n); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
