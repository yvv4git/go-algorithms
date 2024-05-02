package main

import "testing"

func Test_countBinarySubstrings(t *testing.T) {
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
			args: args{s: "00110011"},
			want: 6,
		},
		{
			name: "Test case 2",
			args: args{s: "10101"},
			want: 4,
		},
		{
			name: "Test case 3",
			args: args{s: "000111"},
			want: 3,
		},
		{
			name: "Test case 4",
			args: args{s: "111000"},
			want: 3,
		},
		{
			name: "Test case 5",
			args: args{s: "000000"},
			want: 0,
		},
		{
			name: "Test case 6",
			args: args{s: "111111"},
			want: 0,
		},
		{
			name: "Test case 7",
			args: args{s: "00011"},
			want: 2,
		},
		{
			name: "Test case 8",
			args: args{s: "110011"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBinarySubstrings(tt.args.s); got != tt.want {
				t.Errorf("countBinarySubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
