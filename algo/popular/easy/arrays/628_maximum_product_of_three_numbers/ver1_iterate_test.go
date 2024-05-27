package main

import "testing"

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 24,
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{-1, -2, -3},
			},
			want: -6,
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{-10, -20, 1, 2, 3},
			},
			want: 600,
		},
		{
			name: "Test Case 5",
			args: args{
				nums: []int{10, 20, 30},
			},
			want: 6000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
