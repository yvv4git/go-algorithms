package main

import "testing"

func Test_nthSuperUglyNumber(t *testing.T) {
	type args struct {
		n      int
		primes []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				n:      1,
				primes: []int{2, 3, 5},
			},
			want: 1,
		},
		{
			name: "Test Case 2",
			args: args{
				n:      2,
				primes: []int{2, 3, 5},
			},
			want: 2,
		},
		{
			name: "Test Case 3",
			args: args{
				n:      3,
				primes: []int{2, 3, 5},
			},
			want: 3,
		},
		{
			name: "Test Case 4",
			args: args{
				n:      4,
				primes: []int{2, 3, 5},
			},
			want: 4,
		},
		{
			name: "Test Case 5",
			args: args{
				n:      5,
				primes: []int{2, 3, 5},
			},
			want: 5,
		},
		{
			name: "Test Case 6",
			args: args{
				n:      12,
				primes: []int{2, 3, 5},
			},
			want: 16,
		},
		{
			name: "Test Case 7",
			args: args{
				n:      1,
				primes: []int{2, 7, 13, 19},
			},
			want: 1,
		},
		{
			name: "Test Case 8",
			args: args{
				n:      10,
				primes: []int{2, 7, 13, 19},
			},
			want: 26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthSuperUglyNumber(tt.args.n, tt.args.primes); got != tt.want {
				t.Errorf("nthSuperUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
