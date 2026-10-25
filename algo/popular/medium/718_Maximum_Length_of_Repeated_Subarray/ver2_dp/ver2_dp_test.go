package main

import "testing"

func Test_findLengthV2(t *testing.T) {
	const maxLen = 1000

	maxNums1 := make([]int, maxLen)
	maxNums2 := make([]int, maxLen)
	for i := 0; i < maxLen; i++ {
		maxNums1[i] = i % 101
		maxNums2[i] = i % 101
	}

	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			name:  "Example 1",
			nums1: []int{1, 2, 3, 2, 1},
			nums2: []int{3, 2, 1, 4, 7},
			want:  3,
		},
		{
			name:  "Example 2",
			nums1: []int{0, 0, 0, 0, 0},
			nums2: []int{0, 0, 0, 0, 0},
			want:  5,
		},
		{
			name:  "No common elements",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  0,
		},
		{
			name:  "Single matching elements",
			nums1: []int{5},
			nums2: []int{5},
			want:  1,
		},
		{
			name:  "Single different elements",
			nums1: []int{1},
			nums2: []int{2},
			want:  0,
		},
		{
			name:  "Empty arrays",
			nums1: []int{},
			nums2: []int{},
			want:  0,
		},
		{
			name:  "One empty array",
			nums1: []int{1, 2, 3},
			nums2: []int{},
			want:  0,
		},
		{
			name:  "All zeros",
			nums1: []int{0, 0, 0},
			nums2: []int{0, 0, 0},
			want:  3,
		},
		{
			name:  "Duplicate values",
			nums1: []int{1, 1, 1, 1},
			nums2: []int{1, 1},
			want:  2,
		},
		{
			name:  "Common subarray at the end",
			nums1: []int{1, 2, 3, 4},
			nums2: []int{9, 3, 4},
			want:  2,
		},
		{
			name:  "Max size arrays",
			nums1: maxNums1,
			nums2: maxNums2,
			want:  maxLen,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("findLengthV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
