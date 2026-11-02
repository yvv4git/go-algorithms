package main

import "testing"

func Test_countDistinctV2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		p    int
		want int
	}{
		{name: "example 1", nums: []int{2, 3, 3, 2, 2}, k: 2, p: 2, want: 11},
		{name: "example 2", nums: []int{1, 2, 3, 4}, k: 4, p: 1, want: 10},
		{name: "empty input", nums: []int{}, k: 1, p: 2, want: 0},
		{name: "single suitable", nums: []int{1}, k: 1, p: 2, want: 1},
		{name: "single divisible", nums: []int{2}, k: 1, p: 2, want: 1},
		{name: "duplicate within limit", nums: []int{2, 2}, k: 1, p: 2, want: 1},
		{name: "duplicate next length", nums: []int{2, 2}, k: 2, p: 2, want: 2},
		{name: "fewer divisible than limit", nums: []int{2, 4}, k: 1, p: 2, want: 2},
		{name: "all identical not divisible", nums: []int{3, 3, 3}, k: 1, p: 2, want: 3},
		{name: "all identical divisible", nums: []int{5, 5, 5}, k: 2, p: 5, want: 2},
		{name: "periodic data", nums: []int{1, 2, 1, 2}, k: 1, p: 2, want: 5},
		{name: "all divisible", nums: []int{1, 2, 3, 4}, k: 4, p: 1, want: 10},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if actual := countDistinctV2(testCase.nums, testCase.k, testCase.p); actual != testCase.want {
				t.Fatalf("countDistinctV2() = %d, want %d", actual, testCase.want)
			}
		})
	}
}

func Test_countDistinctV2_maxSize(t *testing.T) {
	nums := make([]int, 200)
	for index := range nums {
		nums[index] = 1
	}

	if actual := countDistinctV2(nums, 1, 2); actual != 200 {
		t.Fatalf("countDistinctV2() = %d, want 200", actual)
	}
}
