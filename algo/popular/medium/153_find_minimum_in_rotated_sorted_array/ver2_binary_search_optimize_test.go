package _53_find_minimum_in_rotated_sorted_array

import "testing"

func Test_findMinV2(t *testing.T) {
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
				nums: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{11, 13, 15, 17},
			},
			want: 11,
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{2, 1},
			},
			want: 1,
		},
		{
			name: "Test Case 5",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinV2(tt.args.nums); got != tt.want {
				t.Errorf("findMinV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
