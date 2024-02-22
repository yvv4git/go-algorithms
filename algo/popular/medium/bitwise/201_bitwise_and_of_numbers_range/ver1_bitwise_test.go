package main

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				left:  5,
				right: 7,
			},
			want: 4,
		},
		{
			name: "Test Case 2",
			args: args{
				left:  0,
				right: 0,
			},
			want: 0,
		},
		{
			name: "Test Case 3",
			args: args{
				left:  1,
				right: 2147483647,
			},
			want: 0,
		},
		{
			name: "Test Case 4",
			args: args{
				left:  10,
				right: 15,
			},
			want: 8,
		},
		{
			name: "Test Case 5",
			args: args{
				left:  3,
				right: 3,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
