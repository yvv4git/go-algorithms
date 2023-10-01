package _45_reverse_vowels_string

import "testing"

func Test_reverseVowelsV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				s: "hello",
			},
			want: "holle",
		},
		{
			name: "CASE-2",
			args: args{
				s: "leetcode",
			},
			want: "leotcede",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowelsV2(tt.args.s); got != tt.want {
				t.Errorf("reverseVowelsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
