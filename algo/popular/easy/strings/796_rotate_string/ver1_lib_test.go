package main

import "testing"

func Test_rotateString(t *testing.T) {
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
			name: "Test Case 1: Strings are identical",
			args: args{
				s:    "abcde",
				goal: "abcde",
			},
			want: true,
		},
		{
			name: "Test Case 2: Strings are rotations",
			args: args{
				s:    "abcde",
				goal: "cdeab",
			},
			want: true,
		},
		{
			name: "Test Case 3: Strings have different lengths",
			args: args{
				s:    "abcde",
				goal: "abcdef",
			},
			want: false,
		},
		{
			name: "Test Case 4: Empty strings",
			args: args{
				s:    "",
				goal: "",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
