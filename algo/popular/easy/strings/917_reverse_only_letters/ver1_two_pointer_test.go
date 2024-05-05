package main

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with all letters",
			args: args{s: "abc"},
			want: "cba",
		},
		{
			name: "Test with letters and non-letters",
			args: args{s: "ab-cd"},
			want: "dc-ba",
		},
		{
			name: "Test with only non-letters",
			args: args{s: "1-2-3"},
			want: "1-2-3",
		},
		{
			name: "Test with empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "Test with single letter",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "Test with single non-letter",
			args: args{s: "-"},
			want: "-",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters(tt.args.s); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
