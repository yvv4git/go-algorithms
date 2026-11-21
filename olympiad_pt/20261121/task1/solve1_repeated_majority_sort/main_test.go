package main

import (
	"reflect"
	"testing"
)

func makeAsker(pi []int, k int, liePattern map[int]bool) func(int, int) bool {
	lies := 0
	calls := 0
	return func(x, y int) bool {
		calls++
		truth := pi[x-1] < pi[y-1]
		if lies < k && liePattern[calls] {
			lies++
			return !truth
		}
		return truth
	}
}

func Test_recoverPermutation(t *testing.T) {
	tests := []struct {
		name       string
		pi         []int
		k          int
		q          int
		liePattern map[int]bool
	}{
		{
			name: "example 1",
			pi:   []int{2, 3, 1},
			k:    1,
			q:    11000,
			liePattern: map[int]bool{
				3: true,
			},
		},
		{
			name:       "no lies sorted",
			pi:         []int{1, 2, 3, 4},
			k:          0,
			q:          11000,
			liePattern: map[int]bool{},
		},
		{
			name:       "no lies reverse",
			pi:         []int{4, 3, 2, 1},
			k:          0,
			q:          11000,
			liePattern: map[int]bool{},
		},
		{
			name: "single lie",
			pi:   []int{3, 1, 2},
			k:    1,
			q:    11000,
			liePattern: map[int]bool{
				2: true,
			},
		},
		{
			name: "two lies",
			pi:   []int{2, 4, 1, 3},
			k:    2,
			q:    11000,
			liePattern: map[int]bool{
				1: true,
				5: true,
			},
		},
		{
			name:       "two elements",
			pi:         []int{2, 1},
			k:          0,
			q:          11000,
			liePattern: map[int]bool{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ask := makeAsker(tt.pi, tt.k, tt.liePattern)
			got := recoverPermutation(len(tt.pi), tt.k, tt.q, ask)
			if !reflect.DeepEqual(got, tt.pi) {
				t.Fatalf("recoverPermutation() = %v, want %v", got, tt.pi)
			}
		})
	}
}

func Test_recoverPermutation_maxSize(t *testing.T) {
	n := 20
	pi := make([]int, n)
	for i := 0; i < n; i++ {
		pi[i] = n - i
	}
	ask := makeAsker(pi, 1, map[int]bool{4: true, 11: true})
	got := recoverPermutation(n, 1, 11000, ask)
	if !reflect.DeepEqual(got, pi) {
		t.Fatalf("recoverPermutation(maxSize) = %v, want %v", got, pi)
	}
}
