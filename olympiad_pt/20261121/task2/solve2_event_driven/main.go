package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"sort"
)

type device struct {
	a         int
	b         int
	t         int
	neighbors []int
}

type input struct {
	n       int
	p       int
	k       int
	x       int
	m       int
	devices []device
}

type dispatch struct {
	u int
	v int
	s int
}

type event struct {
	time   int
	device int
}

type eventHeap []event

func (h eventHeap) Len() int           { return len(h) }
func (h eventHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h eventHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *eventHeap) Push(x any)        { *h = append(*h, x.(event)) }
func (h *eventHeap) Pop() any {
	old := *h
	n := len(old)
	e := old[n-1]
	*h = old[:n-1]
	return e
}

func buildStrategyV2(in input) []dispatch {
	n := in.n
	freeAt := make([]int, n+1)
	for i := 1; i <= n; i++ {
		freeAt[i] = -1
	}

	last := in.x + in.m - 1
	h := &eventHeap{}
	heap.Init(h)
	heap.Push(h, event{time: in.x, device: in.p})

	var result []dispatch
	for h.Len() > 0 {
		e := heap.Pop(h).(event)
		if e.time > last {
			break
		}
		u := e.device

		type candidate struct {
			v      int
			finish int
		}
		cands := make([]candidate, 0, len(in.devices[u].neighbors))
		for _, v := range in.devices[u].neighbors {
			if freeAt[v] > e.time {
				continue
			}
			r := (e.time - in.devices[v].a) % 24
			if r < 0 {
				r += 24
			}
			if r > in.devices[v].b {
				continue
			}
			if r+in.devices[v].t > in.devices[v].b {
				continue
			}
			finish := e.time + in.devices[v].t
			if finish > last {
				continue
			}
			cands = append(cands, candidate{v: v, finish: finish})
		}

		sort.Slice(cands, func(i, j int) bool {
			if cands[i].finish != cands[j].finish {
				return cands[i].finish < cands[j].finish
			}
			return cands[i].v < cands[j].v
		})

		limit := in.k
		if limit > len(cands) {
			limit = len(cands)
		}
		for i := 0; i < limit; i++ {
			v := cands[i].v
			finish := cands[i].finish
			freeAt[v] = finish
			heap.Push(h, event{time: finish, device: v})
			result = append(result, dispatch{u: u, v: v, s: e.time})
		}
	}
	return result
}

func readInput(reader io.Reader) (input, bool) {
	var in input
	if _, err := fmt.Fscan(reader, &in.n, &in.p, &in.k); err != nil {
		return in, false
	}
	if _, err := fmt.Fscan(reader, &in.x, &in.m); err != nil {
		return in, false
	}
	in.devices = make([]device, in.n+1)
	for i := 1; i <= in.n; i++ {
		if _, err := fmt.Fscan(reader, &in.devices[i].a, &in.devices[i].b); err != nil {
			return in, false
		}
	}
	for i := 1; i <= in.n; i++ {
		if _, err := fmt.Fscan(reader, &in.devices[i].t); err != nil {
			return in, false
		}
	}
	for i := 1; i <= in.n; i++ {
		for {
			var v int
			if _, err := fmt.Fscan(reader, &v); err != nil {
				return in, false
			}
			if v == 0 {
				break
			}
			in.devices[i].neighbors = append(in.devices[i].neighbors, v)
		}
	}
	return in, true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	in, ok := readInput(reader)
	if !ok {
		return
	}
	result := buildStrategyV2(in)
	fmt.Fprintln(writer, len(result)+1)
	for _, d := range result {
		fmt.Fprintf(writer, "%d %d %d\n", d.u, d.v, d.s)
	}
}
