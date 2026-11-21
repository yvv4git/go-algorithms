package main

import (
	"reflect"
	"testing"
)

func TestTruthAllN(t *testing.T) {
	for _, n := range []int{2, 3, 4, 5, 10, 50, 100, 200, 500, 1000} {
		for _, k := range []int{0, 1, 5, 10, 30} {
			pi := make([]int, n)
			for i := range pi {
				pi[i] = n - i
			}
			a := &simAsker{pi: pi, k: k, strategy: "truth"}
			q := 11000
			got := recoverPermutationV3(n, k, q, a.ask)
			if a.queries > q {
				t.Errorf("n=%d k=%d overflow %d", n, k, a.queries)
			}
			if !reflect.DeepEqual(got, pi) {
				t.Errorf("n=%d k=%d wrong queries=%d", n, k, a.queries)
			}
		}
	}
}

func TestGuaranteedPathBound(t *testing.T) {
	// All lies at the very end to maximize queries on guaranteed path
	for _, n := range []int{2, 3, 5, 10, 50, 100} {
		for _, k := range []int{0, 1, 2, 5, 10} {
			var cc func(int) int
			cc = func(m int) int {
				if m <= 1 {
					return 0
				}
				return cc(m/2) + cc((m+1)/2) + m - 1
			}
			cMax := cc(n)
			q := (k+1)*cMax + k
			if q > 11000 {
				continue
			}
			pi := make([]int, n)
			for i := range pi {
				pi[i] = i + 1
			}
			liesLeft := k
			queries := 0
			ask := func(x, y int) bool {
				queries++
				truth := pi[x-1] < pi[y-1]
				// lie on second half of budget
				if liesLeft > 0 && queries > q/2 {
					liesLeft--
					return !truth
				}
				return truth
			}
			_ = recoverPermutationV3(n, k, q, ask)
			if queries > q {
				t.Errorf("n=%d k=%d q=%d used=%d", n, k, q, queries)
			}
		}
	}
}
