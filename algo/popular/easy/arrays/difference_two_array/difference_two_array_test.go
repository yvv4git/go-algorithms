package difference_two_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDifferenceWithHashSet(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}

	testCases := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "CASE-1",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4, 6},
			},
			want: [][]int{{1, 3}, {6, 4}},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := FindDifferenceWithHashSet(tt.args.nums1, tt.args.nums2)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestFindDifferenceWithSlices(t *testing.T) {
	var nilSliceInt []int
	type args struct {
		nums1 []int
		nums2 []int
	}

	testCases := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "CASE-1",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4, 6},
			},
			want: [][]int{{1, 3}, {4, 6}},
		},
		{
			name: "CASE-2",
			args: args{
				nums1: []int{1, 2, 3, 3},
				nums2: []int{1, 1, 2, 2},
			},
			want: [][]int{{3}, nilSliceInt},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := FindDifferenceWithSlices(tt.args.nums1, tt.args.nums2)
			assert.EqualValues(t, tt.want, result)
		})
	}
}
