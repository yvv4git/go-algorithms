package main

import "testing"

func Test_isInterleaveV2(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1: Interleaving is possible",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbcbcac",
			},
			want: true,
		},
		{
			name: "Test Case 2: Interleaving is not possible",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbbaccc",
			},
			want: false,
		},
		{
			name: "Test Case 3: One string is empty",
			args: args{
				s1: "",
				s2: "abc",
				s3: "abc",
			},
			want: true,
		},
		{
			name: "Test Case 4: All strings are empty",
			args: args{
				s1: "",
				s2: "",
				s3: "",
			},
			want: true,
		},
		{
			name: "Test Case 5: s3 is shorter than s1 and s2 combined",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbc",
			},
			want: false,
		},
		{
			name: "Test Case 6: s3 is longer than s1 and s2 combined",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbcbcacf",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleaveV2(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleaveV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
