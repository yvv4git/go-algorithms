package is_subsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
