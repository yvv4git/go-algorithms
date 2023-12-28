package _21_maximum_xor_of_two_numbers_in_an_array

import "testing"

func Test_findMaximumXORV2(t *testing.T) {
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
				nums: []int{3, 10, 5, 25, 2, 8},
			},
			want: 28,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{2, 4},
			},
			want: 6,
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{8, 10, 2},
			},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumXORV2(tt.args.nums); got != tt.want {
				t.Errorf("findMaximumXORV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
