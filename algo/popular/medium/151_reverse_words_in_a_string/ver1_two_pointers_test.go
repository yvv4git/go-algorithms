package main

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with single word",
			args: args{s: "Hello"},
			want: "Hello",
		},
		{
			name: "Test with multiple words",
			args: args{s: "Hello World"},
			want: "World Hello",
		},
		{
			name: "Test with multiple words and spaces",
			args: args{s: "  Hello   World  "},
			want: "World Hello",
		},
		{
			name: "Test with multiple words and leading/trailing spaces",
			args: args{s: "  Hello   World"},
			want: "World Hello",
		},
		{
			name: "Test with multiple words and multiple spaces between words",
			args: args{s: "The quick brown fox"},
			want: "fox brown quick The",
		},
		{
			name: "Test with empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "Test with single space",
			args: args{s: " "},
			want: "",
		},
		{
			name: "Test with multiple spaces",
			args: args{s: "   "},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
