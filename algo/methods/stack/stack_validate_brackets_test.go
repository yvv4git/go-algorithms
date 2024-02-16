package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1: Valid parentheses",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "Test case 2: Valid brackets",
			args: args{s: "[]"},
			want: true,
		},
		{
			name: "Test case 3: Valid braces",
			args: args{s: "{}"},
			want: true,
		},
		{
			name: "Test case 4: Valid nested parentheses",
			args: args{s: "(())"},
			want: true,
		},
		{
			name: "Test case 5: Valid nested brackets",
			args: args{s: "{[]}"},
			want: true,
		},
		{
			name: "Test case 6: Valid nested braces",
			args: args{s: "{[()]}"},
			want: true,
		},
		{
			name: "Test case 7: Invalid parentheses",
			args: args{s: "(]"},
			want: false,
		},
		{
			name: "Test case 8: Invalid brackets",
			args: args{s: "{]}"},
			want: false,
		},
		{
			name: "Test case 9: Invalid braces",
			args: args{s: "[}"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
