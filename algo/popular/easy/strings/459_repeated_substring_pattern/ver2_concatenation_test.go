package _59_repeated_substring_pattern

import "testing"

func Test_repeatedSubstringPatternV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{s: "abab"},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{s: "aba"},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{s: "abcabcabcabc"},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{s: "a"},
			want: false,
		},
		{
			name: "Test case 5",
			args: args{s: "aaa"},
			want: true,
		},
		{
			name: "Test case 6",
			args: args{s: "abcabc"},
			want: true,
		},
		{
			name: "Test case 7",
			args: args{s: "abcabcab"},
			want: false,
		},
		{
			name: "Test case 8",
			args: args{s: "ababab"},
			want: true,
		},
		{
			name: "Test case 9",
			args: args{s: "abcdabcd"},
			want: true,
		},
		{
			name: "Test case 10",
			args: args{s: "abcdabcdabcd"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPatternV2(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPatternV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
