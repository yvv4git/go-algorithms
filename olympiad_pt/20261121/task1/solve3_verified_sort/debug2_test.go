package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// Instrument by wrapping - we'll copy the algorithm with logging
func TestDebugPasses(t *testing.T) {
	n, k, q := 100, 20, 11000
	pi := make([]int, n)
	for i := range pi {
		pi[i] = i + 1
	}
	rng := rand.New(rand.NewSource(1))
	rng.Shuffle(n, func(i, j int) { pi[i], pi[j] = pi[j], pi[i] })
	j := &advJury{pi: pi, k: k, liesLeft: k, maxQ: q, mode: "random", rng: rng}

	// Manual reimplementation of the loop with logging
	maxQueries := q
	liesLeft := k
	queriesUsed := 0
	override := map[[2]int]bool{}
	askFresh := func(x, y int) bool {
		if queriesUsed >= maxQueries {
			return x < y
		}
		queriesUsed++
		return j.ask(x, y)
	}
	answer := func(x, y int) bool {
		if v, ok := override[[2]int{x, y}]; ok {
			return v
		}
		if v, ok := override[[2]int{y, x}]; ok {
			return !v
		}
		v := askFresh(x, y)
		override[[2]int{x, y}] = v
		return v
	}
	answerSort := func(x, y int) bool {
		if v, ok := override[[2]int{x, y}]; ok {
			return v
		}
		if v, ok := override[[2]int{y, x}]; ok {
			return !v
		}
		return askFresh(x, y)
	}
	var mergeSort func(arr []int, cmp func(int, int) bool) []int
	mergeSort = func(arr []int, cmp func(int, int) bool) []int {
		if len(arr) <= 1 {
			r := make([]int, len(arr))
			copy(r, arr)
			return r
		}
		mid := len(arr) / 2
		left := mergeSort(arr[:mid], cmp)
		right := mergeSort(arr[mid:], cmp)
		res := make([]int, 0, len(arr))
		i, j2 := 0, 0
		for i < len(left) && j2 < len(right) {
			if cmp(left[i], right[j2]) {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j2])
				j2++
			}
		}
		res = append(res, left[i:]...)
		res = append(res, right[j2:]...)
		return res
	}
	resolve := func(a, b int, seedNo bool) bool {
		yes, no := 0, 0
		if seedNo {
			no = 1
		}
		for {
			if yes > liesLeft {
				liesLeft -= no
				return true
			}
			if no > liesLeft {
				liesLeft -= yes
				return false
			}
			if queriesUsed >= maxQueries {
				if yes == no {
					return a < b
				}
				return yes > no
			}
			queriesUsed++
			if j.ask(a, b) {
				yes++
			} else {
				no++
			}
		}
	}
	order := make([]int, n)
	for i := 0; i < n; i++ {
		order[i] = i + 1
	}
	order = mergeSort(order, answer)
	fmt.Printf("after stage1 queries=%d juryLiesLeft=%d ourLiesLeft=%d\n", queriesUsed, j.liesLeft, liesLeft)

	// count adjacent inversions vs truth
	countInv := func(ord []int) int {
		c := 0
		for i := 0; i+1 < n; i++ {
			if pi[ord[i]-1] > pi[ord[i+1]-1] {
				c++
			}
		}
		return c
	}
	fmt.Printf("stage1 adjacent truth inversions=%d\n", countInv(order))

	minPasses := k + 1
	for pass := 0; ; pass++ {
		changed := false
		nos := 0
		for i := 0; i+1 < n; i++ {
			if queriesUsed >= maxQueries {
				break
			}
			a, b := order[i], order[i+1]
			if askFresh(a, b) {
				if _, ok := override[[2]int{a, b}]; !ok {
					if _, ok2 := override[[2]int{b, a}]; !ok2 {
						override[[2]int{a, b}] = true
					}
				}
				continue
			}
			nos++
			prev, hadPrev := override[[2]int{a, b}]
			if !hadPrev {
				if v, ok := override[[2]int{b, a}]; ok {
					prev, hadPrev = !v, true
				}
			}
			trueAns := resolve(a, b, true)
			if hadPrev && prev != trueAns {
				if liesLeft > 0 {
					liesLeft--
				}
			}
			override[[2]int{a, b}] = trueAns
			if !trueAns {
				changed = true
			}
		}
		order = mergeSort(order, answerSort)
		fmt.Printf("pass=%d nos=%d changed=%v queries=%d ourLiesLeft=%d juryLiesLeft=%d inv=%d\n",
			pass, nos, changed, queriesUsed, liesLeft, j.liesLeft, countInv(order))
		if !changed && pass+1 >= minPasses {
			fmt.Println("break: !changed && minPasses")
			break
		}
		if queriesUsed >= maxQueries {
			fmt.Println("break: budget")
			break
		}
		if pass >= 2*k+n+50 {
			fmt.Println("break: safety")
			break
		}
	}
	// check result
	wrong := 0
	p := make([]int, n)
	for value, pos := range order {
		p[pos-1] = value + 1
	}
	for i := range pi {
		if p[i] != pi[i] {
			wrong++
		}
	}
	fmt.Printf("final wrong=%d queries=%d\n", wrong, queriesUsed)
}
