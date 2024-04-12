package main

import "testing"

func Test_countNumbersWithUniqueDigits(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1: n = 0",
			args: args{n: 0},
			want: 1, // 0
		},
		{
			name: "Test Case 2: n = 1",
			args: args{n: 1},
			want: 10, // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
		},
		{
			name: "Test Case 3: n = 2",
			args: args{n: 2},
			want: 91, // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 - 19, 20 - 29, ..., 90 - 99
		},
		{
			name: "Test Case 4: n = 3",
			args: args{n: 3},
			want: 739, // Similar to n = 2, but for 3 digits
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumbersWithUniqueDigits(tt.args.n); got != tt.want {
				t.Errorf("countNumbersWithUniqueDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
