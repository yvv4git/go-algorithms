package main

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "CASE-2",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "CASE-3",
			args: args{
				haystack: "",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "CASE-4",
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "CASE-5",
			args: args{
				haystack: "aaabbaabb",
				needle:   "bb",
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
