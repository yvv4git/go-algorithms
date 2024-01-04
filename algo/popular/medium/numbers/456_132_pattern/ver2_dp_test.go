package _56_132_pattern

import "testing"

func Test_find132patternV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{3, 1, 4, 2},
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{3, 5, 0, 3, 4},
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{-1, 3, 2, 0},
			},
			want: true,
		},
		{
			name: "Test Case 5",
			args: args{
				nums: []int{1, 0, 1, -4, -3},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132patternV2(tt.args.nums); got != tt.want {
				t.Errorf("find132patternV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
