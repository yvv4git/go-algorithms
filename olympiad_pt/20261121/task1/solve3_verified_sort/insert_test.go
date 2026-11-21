package main

import (
	"reflect"
	"testing"
)

// Test binary-insert repair on nearly sorted arrays with truthful jury
func TestBinaryInsertRepair(t *testing.T) {
	// order of positions such that values are almost sorted with a few inversions
	for _, tc := range []struct {
		name string
		pi   []int // pi[pos-1] = value
		init []int // initial order of positions
	}{
		{"swap2", []int{1, 3, 2, 4, 5}, []int{1, 3, 2, 4, 5}},
		{"far", []int{5, 1, 2, 3, 4}, []int{2, 3, 4, 5, 1}}, // pos5 (val1) at end
		{"rev", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"two", []int{1, 4, 3, 2, 5}, []int{1, 4, 3, 2, 5}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			n := len(tc.pi)
			order := append([]int{}, tc.init...)
			queries := 0
			ask := func(x, y int) bool {
				queries++
				return tc.pi[x-1] < tc.pi[y-1]
			}
			askFresh := func(x, y int) bool {
				return ask(x, y)
			}
			// one repair pass
			for i := 0; i+1 < n; {
				a, b := order[i], order[i+1]
				if askFresh(a, b) {
					i++
					continue
				}
				lo, hi := 0, i+1
				for lo < hi {
					mid := (lo + hi) / 2
					if askFresh(order[mid], b) {
						lo = mid + 1
					} else {
						hi = mid
					}
				}
				for j := i + 1; j > lo; j-- {
					order[j] = order[j-1]
				}
				order[lo] = b
				if lo > 0 {
					i = lo - 1
				} else {
					i = 0
				}
			}
			// verify sorted by value
			for i := 0; i+1 < n; i++ {
				if tc.pi[order[i]-1] > tc.pi[order[i+1]-1] {
					t.Fatalf("still inverted at %d: order=%v pi=%v", i, order, tc.pi)
				}
			}
			// verify permutation property
			seen := map[int]bool{}
			for _, v := range order {
				if seen[v] || v < 1 || v > n {
					t.Fatalf("bad perm %v", order)
				}
				seen[v] = true
			}
			// expected: positions sorted by value
			want := make([]int, n)
			for i := range want {
				want[i] = i + 1
			}
			// sort want by pi
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					if tc.pi[want[i]-1] > tc.pi[want[j]-1] {
						want[i], want[j] = want[j], want[i]
					}
				}
			}
			if !reflect.DeepEqual(order, want) {
				t.Fatalf("got %v want %v", order, want)
			}
			t.Logf("queries=%d", queries)
		})
	}
}
