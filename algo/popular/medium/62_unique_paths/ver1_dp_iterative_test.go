package main

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1: 3x7 grid",
			args: args{m: 3, n: 7},
			want: 28,
		},
		{
			name: "Test case 2: 3x3 grid",
			args: args{m: 3, n: 3},
			want: 6,
		},
		{
			name: "Test case 3: 7x3 grid",
			args: args{m: 7, n: 3},
			want: 28,
		},
		{
			name: "Test case 4: 1x1 grid",
			args: args{m: 1, n: 1},
			want: 1,
		},
		{
			name: "Test case 5: 1x7 grid",
			args: args{m: 1, n: 7},
			want: 1,
		},
		{
			name: "Test case 6: 7x1 grid",
			args: args{m: 7, n: 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
