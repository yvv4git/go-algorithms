package main

import "testing"

func Test_singleNumber(t *testing.T) {
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
				nums: []int{2, 2, 3, 2},
			},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{0, 1, 0, 1, 0, 1, 99},
			},
			want: 99,
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
