package main

import "testing"

func Test_numIdenticalPairs(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 1, 1, 3},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 1, 1, 1},
			},
			want: 6,
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIdenticalPairs(tt.args.nums); got != tt.want {
				t.Errorf("numIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
