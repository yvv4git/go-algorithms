package main

import "testing"

func Test_numSquares(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{n: 12},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{n: 13},
			want: 2,
		},
		{
			name: "Test Case 3",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Test Case 4",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "Test Case 5",
			args: args{n: 4},
			want: 1,
		},
		{
			name: "Test Case 6",
			args: args{n: 18},
			want: 2,
		},
		{
			name: "Test Case 7",
			args: args{n: 100},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.args.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
