package main

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1: Same strings with repeated characters",
			args: args{
				s:    "aa",
				goal: "aa",
			},
			want: true,
		},
		{
			name: "Test Case 2: Different strings with swappable characters",
			args: args{
				s:    "ab",
				goal: "ba",
			},
			want: true,
		},
		{
			name: "Test Case 3: Different strings with non-swappable characters",
			args: args{
				s:    "ab",
				goal: "ab",
			},
			want: false,
		},
		{
			name: "Test Case 4: Different strings with swappable characters",
			args: args{
				s:    "abcd",
				goal: "abdc",
			},
			want: true,
		},
		{
			name: "Test Case 5: Strings of different lengths",
			args: args{
				s:    "abc",
				goal: "abcd",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
