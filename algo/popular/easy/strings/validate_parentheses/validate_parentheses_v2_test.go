package main

import "testing"

func Test_isValidV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "CASE-3",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "CASE-4",
			args: args{
				s: "[}",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidV2(tt.args.s); got != tt.want {
				t.Errorf("isValidV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
