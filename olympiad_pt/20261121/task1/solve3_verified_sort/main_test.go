package main

import (
	"reflect"
	"testing"
)

type simAsker struct {
	pi         []int
	k          int
	liesUsed   int
	queries    int
	strategy   string // "truth", "random", "late", "adaptive"
	rngSeed    int64
	scripted   map[[2]int][]bool
	lieBudget  int
	lieStart   int
}

func (a *simAsker) ask(x, y int) bool {
	a.queries++
	truth := a.pi[x-1] < a.pi[y-1]
	ans := truth
	switch a.strategy {
	case "truth":
		ans = truth
	case "random":
		// lie sporadically while budget allows
		if a.liesUsed < a.k && a.rngSeed%7 == 0 {
			ans = !truth
			a.liesUsed++
		}
		a.rngSeed = a.rngSeed*1103515245 + 12345
	case "adaptive":
		// lie on first k comparisons that involve a specific pattern; simple: lie on k-th distinct pair
		if a.liesUsed < a.k && a.queries <= a.k*3 {
			ans = !truth
			a.liesUsed++
		}
	}
	if a.scripted != nil {
		if answers, ok := a.scripted[[2]int{x, y}]; ok && len(answers) > 0 {
			ans = answers[0]
			a.scripted[[2]int{x, y}] = answers[1:]
			if ans != truth && a.liesUsed < a.k {
				a.liesUsed++
			}
		}
	}
	return ans
}

func Test_recoverPermutationV3_example(t *testing.T) {
	pi := []int{2, 3, 1}
	a := &simAsker{pi: pi, k: 1, strategy: "truth", scripted: map[[2]int][]bool{{2, 3}: {true}}}
	got := recoverPermutationV3(3, 1, 11000, a.ask)
	if !reflect.DeepEqual(got, pi) {
		t.Fatalf("example: got %v want %v queries=%d lies=%d", got, pi, a.queries, a.liesUsed)
	}
}

func Test_recoverPermutationV3_truth(t *testing.T) {
	for _, n := range []int{2, 3, 10, 50, 100, 1000} {
		pi := make([]int, n)
		for i := range pi {
			pi[i] = n - i // reverse
		}
		a := &simAsker{pi: pi, k: 0, strategy: "truth"}
		got := recoverPermutationV3(n, 0, 11000, a.ask)
		if !reflect.DeepEqual(got, pi) {
			t.Fatalf("truth n=%d: got %v want %v queries=%d", n, got, pi, a.queries)
		}
	}
}

func Test_recoverPermutationV3_small_exhaustive(t *testing.T) {
	// Все перестановки для n=4, k=0..2, стратегия truth (без лжи — должно быть точно)
	perms := permute(4)
	for _, pi := range perms {
		for _, k := range []int{0, 1, 2} {
			a := &simAsker{pi: pi, k: k, strategy: "truth"}
			got := recoverPermutationV3(4, k, 10000, a.ask)
			if !reflect.DeepEqual(got, pi) {
				t.Fatalf("truth n=4 k=%d pi=%v: got %v queries=%d", k, pi, got, a.queries)
			}
		}
	}
}

func Test_recoverPermutationV3_query_budget(t *testing.T) {
	// Проверяем, что запросов не превышает Q.
	for _, n := range []int{100, 500, 1000} {
		for _, k := range []int{0, 1, 5, 10, 30} {
			pi := make([]int, n)
			for i := range pi {
				pi[i] = i + 1
			}
			// стратегия truth: без лжи
			a := &simAsker{pi: pi, k: k, strategy: "truth"}
			got := recoverPermutationV3(n, k, 11000, a.ask)
			if a.queries > 11000 {
				t.Fatalf("n=%d k=%d queries=%d > Q", n, k, a.queries)
			}
			if !reflect.DeepEqual(got, pi) {
				t.Fatalf("n=%d k=%d truth: wrong answer queries=%d", n, k, a.queries)
			}
		}
	}
}

func permute(n int) [][]int {
	var res [][]int
	var rec func(arr []int, start int)
	rec = func(arr []int, start int) {
		if start == n {
			out := make([]int, n)
			copy(out, arr)
			res = append(res, out)
			return
		}
		for i := start; i < n; i++ {
			arr[start], arr[i] = arr[i], arr[start]
			rec(arr, start+1)
			arr[start], arr[i] = arr[i], arr[start]
		}
	}
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}
	rec(arr, 0)
	return res
}
