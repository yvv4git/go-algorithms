package main

import (
	"math/rand"
	"testing"
)

// adversarial jury: knows pi, adaptive lies to maximize damage
type advJury struct {
	pi       []int
	k        int
	liesLeft int
	queries  int
	maxQ     int
	// strategy:
	//  "corrupt": lie when truthful answer would preserve sortedness of our
	//             emerging order — approximated by lying on comparisons that
	//             flip merge outcomes, saving lies for verification phase
	//  "random":  lie randomly
	//  "late":    save lies for later queries
	mode string
	rng  *rand.Rand
	// track which pairs we've lied on to hide inversions in final checks
	lieOnYes map[int]bool // for late mode: lie to say YES when truth is NO
}

func (j *advJury) ask(x, y int) bool {
	j.queries++
	truth := j.pi[x-1] < j.pi[y-1]
	if j.liesLeft <= 0 {
		return truth
	}
	switch j.mode {
	case "random":
		if j.rng.Intn(3) == 0 {
			j.liesLeft--
			return !truth
		}
	case "late":
		// lie in last 40% of budget when it hides an inversion (truth NO -> say YES)
		if j.queries > j.maxQ*3/5 && !truth {
			j.liesLeft--
			return true // hide inversion
		}
		// also lie early on some to corrupt sort
		if j.queries <= j.k*5 && j.rng.Intn(2) == 0 {
			j.liesLeft--
			return !truth
		}
	case "corrupt":
		// Lie when truth would give a "correct-looking" answer that we want to break.
		// Simple: lie on ~half of comparisons early to corrupt sort, rest to hide.
		phase1 := j.queries <= 9000
		if phase1 {
			// lie to corrupt: always lie with prob to spread lies across sort
			if j.rng.Intn(4) == 0 {
				j.liesLeft--
				return !truth
			}
		} else {
			// verification phase: hide inversions
			if !truth {
				j.liesLeft--
				return true
			}
		}
	case "targeted":
		// Save all lies for verification: sort phase truth, then hide inversions
		if !truth {
			// only lie during verification (query > C approx) — use late-like
			if j.queries > 8500 {
				j.liesLeft--
				return true
			}
		}
	}
	return truth
}

func runAdv(n, k, q int, mode string, seed int64) (ok bool, queries int) {
	pi := make([]int, n)
	for i := range pi {
		pi[i] = i + 1
	}
	rng := rand.New(rand.NewSource(seed))
	rng.Shuffle(n, func(i, j int) { pi[i], pi[j] = pi[j], pi[i] })
	j := &advJury{pi: pi, k: k, liesLeft: k, maxQ: q, mode: mode, rng: rng}
	got := recoverPermutationV3(n, k, q, j.ask)
	for i := range pi {
		if got[i] != pi[i] {
			return false, j.queries
		}
	}
	return true, j.queries
}

func TestStressAdversarial(t *testing.T) {
	modes := []string{"random", "late", "corrupt", "targeted"}
	type cfg struct{ n, k, q int }
	cfgs := []cfg{
		{10, 1, 11000},
		{10, 5, 11000},
		{50, 3, 11000},
		{50, 10, 11000},
		{100, 5, 11000},
		{100, 20, 11000},
		{200, 10, 11000},
		{500, 10, 11000},
		{500, 30, 11000},
		{1000, 1, 11000},
		{1000, 5, 11000},
		{1000, 30, 11000},
		{3, 1, 11000},
		{2, 30, 11000},
		{2, 30, 50},
		{2, 30, 61},
		{2, 30, 40},
	}
	fails := 0
	total := 0
	for _, c := range cfgs {
		for _, m := range modes {
			for seed := int64(1); seed <= 20; seed++ {
				total++
				ok, q := runAdv(c.n, c.k, c.q, m, seed)
				if !ok {
					fails++
					if fails <= 30 {
						t.Errorf("FAIL n=%d k=%d q=%d mode=%s seed=%d queries=%d", c.n, c.k, c.q, m, seed, q)
					}
				}
				if q > c.q {
					t.Errorf("OVERFLOW n=%d k=%d q=%d mode=%s seed=%d used=%d", c.n, c.k, c.q, m, seed, q)
				}
			}
		}
	}
	t.Logf("total=%d fails=%d", total, fails)
}

func TestStressQueryOverflowBound(t *testing.T) {
	// Check solve2 path never exceeds Q when (k+1)*cMax <= Q
	// by using truthful jury with adversarial lie placement at end
	for _, k := range []int{0, 1, 5, 10, 30} {
		for _, n := range []int{2, 3, 5, 10, 50, 100, 200, 500, 1000} {
			var countComp func(int) int
			countComp = func(m int) int {
				if m <= 1 {
					return 0
				}
				return countComp(m/2) + countComp((m+1)/2) + m - 1
			}
			cMax := countComp(n)
			// use Q = (k+1)*cMax exactly (old condition) to see if overflow
			q := (k + 1) * cMax
			if q == 0 || q > 11000 {
				continue
			}
			// adversarial: lie on last queries
			pi := make([]int, n)
			for i := range pi {
				pi[i] = n - i // reverse
			}
			liesLeft := k
			queries := 0
			ask := func(x, y int) bool {
				queries++
				truth := pi[x-1] < pi[y-1]
				// lie whenever possible on later queries to maximize cost
				if liesLeft > 0 && queries > q/2 {
					liesLeft--
					return !truth
				}
				return truth
			}
			got := recoverPermutationV3(n, k, q, ask)
			if queries > q {
				t.Errorf("overflow n=%d k=%d q=%d used=%d cMax=%d", n, k, q, queries, cMax)
			}
			_ = got
		}
	}
}
