package main

import "testing"

func Test_largestPerimeter(t *testing.T) {
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
			args: args{nums: []int{2, 1, 2}},
			want: 5,
		},
		{
			name: "Test Case 2",
			args: args{nums: []int{1, 2, 1}},
			want: 0,
		},
		{
			name: "Test Case 3",
			args: args{nums: []int{3, 2, 3, 4}},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.nums); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
