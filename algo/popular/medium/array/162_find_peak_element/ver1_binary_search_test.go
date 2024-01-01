package _62_find_peak_element

import "testing"

func Test_findPeakElementV1(t *testing.T) {
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
			want: 2,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			want: 5,
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: 0,
		},
		{
			name: "Test Case 5",
			args: args{
				nums: []int{1, 2, 3, 2, 1},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElementV1(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElementV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
