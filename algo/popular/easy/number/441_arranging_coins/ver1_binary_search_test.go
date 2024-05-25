package main

import "testing"

func Test_arrangeCoins(t *testing.T) {
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
			args: args{n: 5},
			want: 2,
		},
		{
			name: "Test Case 2",
			args: args{n: 8},
			want: 3,
		},
		{
			name: "Test Case 3",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "Test Case 4",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Test Case 5",
			args: args{n: 10},
			want: 4,
		},
		{
			name: "Test Case 6",
			args: args{n: 1804289383},
			want: 60070,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoins(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
