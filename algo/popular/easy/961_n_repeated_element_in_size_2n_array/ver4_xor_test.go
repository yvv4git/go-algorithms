package main

import "testing"

func Test_repeatedNTimesV4(t *testing.T) {
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
				nums: []int{1, 2, 3, 3},
			},
			want: 3,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{2, 1, 2, 5, 3, 2},
			},
			want: 2,
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{5, 1, 5, 2, 5, 3, 5, 4},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedNTimesV4(tt.args.nums); got != tt.want {
				t.Errorf("repeatedNTimesV4() = %v, want %v", got, tt.want)
			}
		})
	}
}