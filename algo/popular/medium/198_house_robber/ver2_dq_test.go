package main

import "testing"

func Test_robV2(t *testing.T) {
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
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
			},
			want: 12,
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{5, 3, 4, 11, 2},
			},
			want: 16,
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "Test Case 5",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robV2(tt.args.nums); got != tt.want {
				t.Errorf("robV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
