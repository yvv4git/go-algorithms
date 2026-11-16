package qos

import (
	"math/rand"
	"sync"
	"sync/atomic"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

//
// Options
//

type Option[T Integer] func(*WeightedQoS[T])

func WithPriorityRange[T Integer](min, max T) Option[T] {
	return func(q *WeightedQoS[T]) {
		q.minPriority = min
		q.maxPriority = max
	}
}

//
// WeightedQoS
//

const (
	defaultMinPriority = 0
	defaultMaxPriority = 10
)

type WeightedQoS[T Integer] struct {
	mu          sync.Mutex
	minPriority T
	maxPriority T
	totalWeight int
	priorities  []T
	cumWeights  []int
	counter     atomic.Int64
}

func New[T Integer](opts ...Option[T]) *WeightedQoS[T] {
	q := &WeightedQoS[T]{
		minPriority: defaultMinPriority,
		maxPriority: defaultMaxPriority,
	}

	for _, opt := range opts {
		opt(q)
	}

	q.init()

	return q
}

func (q *WeightedQoS[T]) init() {
	count := int(q.maxPriority - q.minPriority + 1)
	q.cumWeights = make([]int, count)
	q.totalWeight = 0

	for i := range count {
		priority := q.minPriority + T(i)
		w := int(priority - q.minPriority + 1)
		q.totalWeight += w
		q.cumWeights[i] = q.totalWeight
	}

	for i := count - 1; i >= 0; i-- {
		priority := q.minPriority + T(i)
		w := int(priority - q.minPriority + 1)

		for range w {
			q.priorities = append(q.priorities, priority)
		}
	}
}

func (q *WeightedQoS[T]) TotalWeight() int {
	return q.totalWeight
}

func (q *WeightedQoS[T]) Pick() T {
	return q.PickRandom()
}

func (q *WeightedQoS[T]) PickRandom() T {
	q.mu.Lock()
	r := rand.Intn(q.totalWeight)
	q.mu.Unlock()
	for i, cw := range q.cumWeights {
		if r < cw {
			return q.minPriority + T(i)
		}
	}
	return q.maxPriority
}

func (q *WeightedQoS[T]) PickCyclic() T {
	idx := int(q.counter.Add(1) - 1)
	return q.priorities[idx%len(q.priorities)]
}

func (q *WeightedQoS[T]) SkipToPriority(priority T) {
	if priority < q.minPriority {
		priority = q.minPriority
	}

	if priority > q.maxPriority {
		priority = q.maxPriority
	}

	firstIdx := 0
	for p := q.maxPriority; p > priority; p-- {
		firstIdx += int(p - q.minPriority + 1)
	}

	q.counter.Store(int64(firstIdx))
}
