package main

import "testing"

func Test_majorityElementSort(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
		{
			name: "CASE-4",
			args: args{
				nums: []int{1, 1, 2, 2, 1, 2, 2},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementSort(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
