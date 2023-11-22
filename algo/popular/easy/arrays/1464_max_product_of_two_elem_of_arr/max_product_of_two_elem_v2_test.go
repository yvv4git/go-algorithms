package _464_max_product_of_two_elem_of_arr

import "testing"

func Test_maxProductV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{3, 4, 5, 2},
			},
			want: 12,
			desc: `If you choose the indices i=1 and j=2 (indexed from 0), you will get the maximum value, that is, (nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12.`,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{1, 5, 4, 5},
			},
			want: 16,
			desc: `Choosing the indices i=1 and j=3 (indexed from 0), you will get the maximum value of (5-1)*(5-1) = 16.`,
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{3, 7},
			},
			want: 12,
			desc: `Choosing the indices i=1 and j=3 (indexed from 0), you will get the maximum value of (5-1)*(5-1) = 16.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductV2(tt.args.nums); got != tt.want {
				t.Errorf("maxProductV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
