package main

import (
	"reflect"
	"testing"
)

type scriptAsker struct {
	pi     []int
	k      int
	lies   int
	calls  int
	script map[[2]int][]bool
}

func (a *scriptAsker) ask(x, y int) bool {
	a.calls++
	key := [2]int{x, y}
	truth := a.pi[x-1] < a.pi[y-1]
	if answers, ok := a.script[key]; ok && len(answers) > 0 {
		ans := answers[0]
		a.script[key] = answers[1:]
		if ans != truth && a.lies < a.k {
			a.lies++
			return ans
		}
		return truth
	}
	return truth
}

func Test_recoverPermutationV2(t *testing.T) {
	tests := []struct {
		name   string
		pi     []int
		k      int
		q      int
		script map[[2]int][]bool
	}{
		{
			name: "example 1",
			pi:   []int{2, 3, 1},
			k:    1,
			q:    11000,
			script: map[[2]int][]bool{
				{2, 3}: {true, false},
			},
		},
		{
			name:   "no lies sorted",
			pi:     []int{1, 2, 3, 4},
			k:      0,
			q:      11000,
			script: map[[2]int][]bool{},
		},
		{
			name:   "no lies reverse",
			pi:     []int{4, 3, 2, 1},
			k:      0,
			q:      11000,
			script: map[[2]int][]bool{},
		},
		{
			name: "single lie",
			pi:   []int{3, 1, 2},
			k:    1,
			q:    11000,
			script: map[[2]int][]bool{
				{1, 2}: {true, false},
			},
		},
		{
			name: "two lies same pair",
			pi:   []int{1, 3, 2},
			k:    2,
			q:    11000,
			script: map[[2]int][]bool{
				{2, 3}: {true, true, false},
			},
		},
		{
			name:   "two elements",
			pi:     []int{2, 1},
			k:      0,
			q:      11000,
			script: map[[2]int][]bool{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &scriptAsker{pi: tt.pi, k: tt.k, script: tt.script}
			got := recoverPermutationV2(len(tt.pi), tt.k, tt.q, a.ask)
			if !reflect.DeepEqual(got, tt.pi) {
				t.Fatalf("recoverPermutationV2() = %v, want %v", got, tt.pi)
			}
			if a.lies > tt.k {
				t.Fatalf("lies = %d, want <= %d", a.lies, tt.k)
			}
		})
	}
}

func Test_recoverPermutationV2_maxSize(t *testing.T) {
	n := 20
	pi := make([]int, n)
	for i := 0; i < n; i++ {
		pi[i] = n - i
	}
	a := &scriptAsker{pi: pi, k: 1, script: map[[2]int][]bool{{1, 2}: {true, false}}}
	got := recoverPermutationV2(n, 1, 11000, a.ask)
	if !reflect.DeepEqual(got, pi) {
		t.Fatalf("recoverPermutationV2(maxSize) = %v, want %v", got, pi)
	}
}
