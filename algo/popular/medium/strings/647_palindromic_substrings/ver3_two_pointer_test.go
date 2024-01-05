package _47_palindromic_substrings

import "testing"

func Test_countSubstringsV3(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{s: "abc"},
			want: 3,
		},
		{
			name: "Test case 2",
			args: args{s: "aaa"},
			want: 6,
		},
		{
			name: "Test case 3",
			args: args{s: "abccba"},
			want: 9,
		},
		{
			name: "Test case 4",
			args: args{s: "racecar"},
			want: 10,
		},
		{
			name: "Test case 5",
			args: args{s: "a"},
			want: 1,
		},
		{
			name: "Test case 6",
			args: args{s: ""},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstringsV3(tt.args.s); got != tt.want {
				t.Errorf("countSubstringsV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
