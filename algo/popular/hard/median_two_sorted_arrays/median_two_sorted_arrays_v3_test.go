package median_two_sorted_arrays

import "testing"

func Test_findMedianSortedArraysV3(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "CASE-1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "CASE-1",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArraysV3(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArraysV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
