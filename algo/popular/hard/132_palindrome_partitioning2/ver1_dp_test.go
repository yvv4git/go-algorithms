package main

import "testing"

func Test_minCut(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{s: "aab"},
			want: 1,
		},
		{
			name: "Test Case 2",
			args: args{s: "a"},
			want: 0,
		},
		{
			name: "Test Case 3",
			args: args{s: "ab"},
			want: 1,
		},
		{
			name: "Test Case 4",
			args: args{s: "aba"},
			want: 0,
		},
		{
			name: "Test Case 5",
			args: args{s: "ababababababababababababcbabababababababababababa"},
			want: 0,
		},
		{
			name: "Test Case 6",
			args: args{s: "racecar"},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
