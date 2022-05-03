package main

import "testing"

func Test_calcSumOfSquares(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				n: 5,
			},
			want: 55,
		},
		{
			name: "CASE-2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "CASE-3",
			args: args{
				n: 0,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := calcSumOfSquares(tt.args.n); result != tt.want {
				t.Errorf("calcSumOfSquares() = %v, want %v", result, tt.want)
			}
		})
	}
}
