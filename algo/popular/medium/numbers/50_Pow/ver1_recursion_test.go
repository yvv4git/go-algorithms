package main

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Case 1: Positive base and positive exponent",
			args: args{x: 2.0, n: 3},
			want: 8.0,
		},
		{
			name: "Test Case 2: Positive base and zero exponent",
			args: args{x: 2.0, n: 0},
			want: 1.0,
		},
		{
			name: "Test Case 3: Positive base and negative exponent",
			args: args{x: 2.0, n: -3},
			want: 0.125,
		},
		{
			name: "Test Case 4: Zero base and positive exponent",
			args: args{x: 0.0, n: 3},
			want: 0.0,
		},
		{
			name: "Test Case 5: Zero base and zero exponent",
			args: args{x: 0.0, n: 0},
			want: 1.0, // According to mathematical convention 0^0 is 1
		},
		{
			name: "Test Case 6: Negative base and even exponent",
			args: args{x: -2.0, n: 4},
			want: 16.0,
		},
		{
			name: "Test Case 7: Negative base and odd exponent",
			args: args{x: -2.0, n: 3},
			want: -8.0,
		},
		{
			name: "Test Case 8: Negative base and negative exponent",
			args: args{x: -2.0, n: -3},
			want: -0.125,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
