package qos

import (
	"sync"
	"testing"
)

func TestNew_Defaults(t *testing.T) {
	q := New[int]()

	if q.minPriority != 0 {
		t.Errorf("minPriority: expected 0, got %d", q.minPriority)
	}
	if q.maxPriority != 10 {
		t.Errorf("maxPriority: expected 10, got %d", q.maxPriority)
	}
	if q.totalWeight != 66 {
		t.Errorf("totalWeight: expected 66, got %d", q.totalWeight)
	}
}

func TestNew_CustomRange(t *testing.T) {
	q := New(WithPriorityRange(1, 5))

	if q.minPriority != 1 {
		t.Errorf("minPriority: expected 1, got %d", q.minPriority)
	}
	if q.maxPriority != 5 {
		t.Errorf("maxPriority: expected 5, got %d", q.maxPriority)
	}

	expectedTotal := 1 + 2 + 3 + 4 + 5
	if q.totalWeight != expectedTotal {
		t.Errorf("totalWeight: expected %d, got %d", expectedTotal, q.totalWeight)
	}
}

func TestNew_CustomRangeInt32(t *testing.T) {
	q := New(WithPriorityRange[int32](1, 5))

	if q.minPriority != 1 {
		t.Errorf("minPriority: expected 1, got %d", q.minPriority)
	}
	if q.maxPriority != 5 {
		t.Errorf("maxPriority: expected 5, got %d", q.maxPriority)
	}

	expectedTotal := 1 + 2 + 3 + 4 + 5
	if q.totalWeight != expectedTotal {
		t.Errorf("totalWeight: expected %d, got %d", expectedTotal, q.totalWeight)
	}
}

func TestNew_SinglePriority(t *testing.T) {
	q := New(WithPriorityRange(3, 3))

	if q.totalWeight != 1 {
		t.Errorf("totalWeight: expected 1, got %d", q.totalWeight)
	}

	for range 10 {
		if p := q.PickRandom(); p != 3 {
			t.Errorf("PickRandom: expected 3, got %d", p)
		}
	}

	for range 10 {
		if p := q.PickCyclic(); p != 3 {
			t.Errorf("PickCyclic: expected 3, got %d", p)
		}
	}
}

func TestTotalWeight(t *testing.T) {
	q := New[int]()
	if q.TotalWeight() != 66 {
		t.Errorf("TotalWeight: expected 66, got %d", q.TotalWeight())
	}
}

func TestPickRandom_Range(t *testing.T) {
	q := New[int]()
	for range 1000 {
		p := q.PickRandom()
		if p < 0 || p > 10 {
			t.Errorf("PickRandom: out of range [0, 10], got %d", p)
		}
	}
}

func TestPickRandom_RangeCustom(t *testing.T) {
	q := New(WithPriorityRange(3, 7))
	for range 1000 {
		p := q.PickRandom()
		if p < 3 || p > 7 {
			t.Errorf("PickRandom: out of range [3, 7], got %d", p)
		}
	}
}

func TestPickRandom_Distribution(t *testing.T) {
	q := New[int]()
	samples := 20000
	counts := make([]int, 11)

	for range samples {
		counts[q.PickRandom()]++
	}

	for p := 0; p <= 10; p++ {
		expected := float64(samples) * float64(p+1) / float64(q.totalWeight)
		got := float64(counts[p])
		ratio := got / expected

		if ratio < 0.5 || ratio > 1.5 {
			t.Errorf("priority %d: expected ~%.0f samples, got %d (ratio %.2f)", p, expected, counts[p], ratio)
		}
	}
}

func TestPickCyclic_ExactCycle(t *testing.T) {
	q := New[int]()
	n := q.totalWeight
	counts := make([]int, 11)

	for range n {
		counts[q.PickCyclic()]++
	}

	for p := 0; p <= 10; p++ {
		expected := p + 1
		if counts[p] != expected {
			t.Errorf("priority %d: expected %d occurrences in full cycle, got %d", p, expected, counts[p])
		}
	}
}

func TestPickCyclic_Range(t *testing.T) {
	q := New[int]()
	for range 1000 {
		p := q.PickCyclic()
		if p < 0 || p > 10 {
			t.Errorf("PickCyclic: out of range [0, 10], got %d", p)
		}
	}
}

func TestPickCyclic_Repeatability(t *testing.T) {
	q := New[int]()
	n := q.totalWeight

	first := make([]int, n)
	for i := range n {
		first[i] = q.PickCyclic()
	}

	second := make([]int, n)
	for i := range n {
		second[i] = q.PickCyclic()
	}

	for i := range n {
		if first[i] != second[i] {
			t.Errorf("cycle mismatch at position %d: first=%d, second=%d", i, first[i], second[i])
		}
	}
}

func TestSkipToPriority(t *testing.T) {
	t.Run("skip from start to middle priority", func(t *testing.T) {
		q := New[int]()
		q.SkipToPriority(2)
		p := q.PickCyclic()
		if p != 2 {
			t.Errorf("expected 2 after SkipToPriority(2), got %d", p)
		}
	})

	t.Run("continue cycle after skip", func(t *testing.T) {
		q := New[int]()
		q.SkipToPriority(2)

		expected := []int{2, 2, 2, 1, 1, 0}
		for i, exp := range expected {
			p := q.PickCyclic()
			if p != exp {
				t.Errorf("position %d after SkipToPriority(2): expected %d, got %d", i, exp, p)
			}
		}
	})

	t.Run("skip to max priority", func(t *testing.T) {
		q := New[int]()
		q.SkipToPriority(10)
		p := q.PickCyclic()
		if p != 10 {
			t.Errorf("expected 10 after SkipToPriority(10), got %d", p)
		}
	})

	t.Run("skip below min clamps to min", func(t *testing.T) {
		q := New(WithPriorityRange(5, 10))
		q.SkipToPriority(0)
		p := q.PickCyclic()
		if p != 5 {
			t.Errorf("expected 5 (clamped to min) after SkipToPriority(0), got %d", p)
		}
	})

	t.Run("skip above max clamps to max", func(t *testing.T) {
		q := New(WithPriorityRange(5, 10))
		q.SkipToPriority(99)
		p := q.PickCyclic()
		if p != 10 {
			t.Errorf("expected 10 after SkipToPriority(99), got %d", p)
		}
	})

	t.Run("full cycle after skip covers all priorities", func(t *testing.T) {
		q := New[int]()
		q.SkipToPriority(5)
		n := q.totalWeight
		counts := make([]int, 11)

		for range n {
			counts[q.PickCyclic()]++
		}

		for p := 0; p <= 10; p++ {
			expected := p + 1
			if counts[p] != expected {
				t.Errorf("priority %d: expected %d occurrences in full cycle, got %d", p, expected, counts[p])
			}
		}
	})
}

func TestPick_DelegatesToRandom(t *testing.T) {
	q := New[int]()
	for range 100 {
		p := q.Pick()
		if p < 0 || p > 10 {
			t.Errorf("Pick: out of range [0, 10], got %d", p)
		}
	}
}

func TestConcurrent_PickRandom(t *testing.T) {
	q := New[int]()
	var wg sync.WaitGroup

	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 500 {
				p := q.PickRandom()
				if p < 0 || p > 10 {
					t.Errorf("PickRandom: out of range [0, 10], got %d", p)
				}
			}
		}()
	}

	wg.Wait()
}

func TestConcurrent_PickCyclic(t *testing.T) {
	q := New[int]()
	var wg sync.WaitGroup

	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 500 {
				p := q.PickCyclic()
				if p < 0 || p > 10 {
					t.Errorf("PickCyclic: out of range [0, 10], got %d", p)
				}
			}
		}()
	}

	wg.Wait()
}

func TestConcurrent_Mixed(t *testing.T) {
	q := New[int]()
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 500 {
				q.PickRandom()
				q.PickCyclic()
			}
		}()
	}

	wg.Wait()
}
