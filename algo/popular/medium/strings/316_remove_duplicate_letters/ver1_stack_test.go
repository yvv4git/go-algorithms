package _16_remove_duplicate_letters

import "testing"

func Test_removeDuplicateLettersV1(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{s: "bcabc"},
			want: "abc",
		},
		{
			name: "Test case 2",
			args: args{s: "cbacdcbc"},
			want: "acdb",
		},
		{
			name: "Test case 3",
			args: args{s: "abacb"},
			want: "abc",
		},
		{
			name: "Test case 4",
			args: args{s: "bbcaac"},
			want: "bac",
		},
		{
			name: "Test case 5",
			args: args{s: "abcdefghijklmnopqrstuvwxyz"},
			want: "abcdefghijklmnopqrstuvwxyz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLettersV1(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicateLettersV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
