package main

import "testing"

func Test_largestNumberV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{nums: []int{10, 2}},
			want: "210",
		},
		{
			name: "Test Case 2",
			args: args{nums: []int{3, 30, 34, 5, 9}},
			want: "9534330",
		},
		{
			name: "Test Case 3",
			args: args{nums: []int{0, 0}},
			want: "0",
		},
		{
			name: "Test Case 4",
			args: args{nums: []int{824, 8247}},
			want: "8248247",
		},
		{
			name: "Test Case 5",
			args: args{nums: []int{121, 12}},
			want: "12121",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumberV2(tt.args.nums); got != tt.want {
				t.Errorf("largestNumberV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
