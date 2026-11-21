package main

import (
	"fmt"
	"sort"
	"testing"
)

func sampleInput() input {
	return input{
		n: 4,
		p: 1,
		k: 1,
		x: 10,
		m: 13,
		devices: []device{
			{},
			{a: 10, b: 9, t: 2, neighbors: []int{2}},
			{a: 6, b: 11, t: 2, neighbors: []int{1, 3}},
			{a: 9, b: 4, t: 1, neighbors: []int{2, 4}},
			{a: 6, b: 4, t: 1, neighbors: []int{3}},
		},
	}
}

func validate(in input, res []dispatch) error {
	type groupKey struct {
		u int
		s int
	}
	groups := make(map[groupKey]int)
	seenTarget := make(map[[2]int]bool)

	for _, d := range res {
		if d.u < 1 || d.u > in.n || d.v < 1 || d.v > in.n {
			return fmt.Errorf("device out of range: %d %d", d.u, d.v)
		}
		if d.u == d.v {
			return fmt.Errorf("self dispatch on device %d", d.u)
		}
		found := false
		for _, nb := range in.devices[d.u].neighbors {
			if nb == d.v {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("device %d is not a neighbor of %d", d.v, d.u)
		}
		if d.s < in.x || d.s > in.x+in.m-1 {
			return fmt.Errorf("dispatch time %d out of range", d.s)
		}
		groups[groupKey{d.u, d.s}]++
		if seenTarget[[2]int{d.v, d.s}] {
			return fmt.Errorf("two copies to device %d at time %d", d.v, d.s)
		}
		seenTarget[[2]int{d.v, d.s}] = true
	}

	for k, cnt := range groups {
		if cnt > in.k {
			return fmt.Errorf("device %d sent %d copies at time %d, limit %d", k.u, cnt, k.s, in.k)
		}
	}

	sorted := make([]dispatch, len(res))
	copy(sorted, res)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].s != sorted[j].s {
			return sorted[i].s < sorted[j].s
		}
		if sorted[i].u != sorted[j].u {
			return sorted[i].u < sorted[j].u
		}
		return sorted[i].v < sorted[j].v
	})

	active := map[[2]int]bool{{in.p, in.x}: true}
	freeAt := make([]int, in.n+1)
	for i := 1; i <= in.n; i++ {
		freeAt[i] = -1
	}

	for _, d := range sorted {
		if !active[[2]int{d.u, d.s}] {
			return fmt.Errorf("no active copy on device %d at time %d", d.u, d.s)
		}
		if freeAt[d.v] > d.s {
			return fmt.Errorf("device %d is busy at time %d", d.v, d.s)
		}
		r := (d.s - in.devices[d.v].a) % 24
		if r < 0 {
			r += 24
		}
		if r > in.devices[d.v].b {
			return fmt.Errorf("device %d is off at time %d", d.v, d.s)
		}
		if r+in.devices[d.v].t > in.devices[d.v].b {
			return fmt.Errorf("unpacking on device %d at time %d does not fit", d.v, d.s)
		}
		finish := d.s + in.devices[d.v].t
		if finish > in.x+in.m-1 {
			return fmt.Errorf("unpacking on device %d finishes too late", d.v)
		}
		freeAt[d.v] = finish
		active[[2]int{d.v, finish}] = true
	}
	return nil
}

func Test_buildStrategyV2(t *testing.T) {
	tests := []struct {
		name          string
		in            input
		minAnchorings int
	}{
		{
			name:          "example 1",
			in:            sampleInput(),
			minAnchorings: 5,
		},
		{
			name: "no neighbors",
			in: input{
				n: 1, p: 1, k: 1, x: 0, m: 10,
				devices: []device{
					{},
					{a: 0, b: 23, t: 1},
				},
			},
			minAnchorings: 1,
		},
		{
			name: "k zero",
			in: input{
				n: 2, p: 1, k: 0, x: 0, m: 10,
				devices: []device{
					{},
					{a: 0, b: 23, t: 1, neighbors: []int{2}},
					{a: 0, b: 23, t: 1, neighbors: []int{1}},
				},
			},
			minAnchorings: 1,
		},
		{
			name: "neighbor off at start",
			in: input{
				n: 2, p: 1, k: 1, x: 0, m: 10,
				devices: []device{
					{},
					{a: 0, b: 23, t: 1, neighbors: []int{2}},
					{a: 2, b: 1, t: 1, neighbors: []int{1}},
				},
			},
			minAnchorings: 1,
		},
		{
			name: "exact closing",
			in: input{
				n: 2, p: 1, k: 1, x: 0, m: 5,
				devices: []device{
					{},
					{a: 0, b: 23, t: 1, neighbors: []int{2}},
					{a: 0, b: 1, t: 1, neighbors: []int{1}},
				},
			},
			minAnchorings: 3,
		},
		{
			name: "closing before finish",
			in: input{
				n: 2, p: 1, k: 1, x: 0, m: 5,
				devices: []device{
					{},
					{a: 0, b: 23, t: 1, neighbors: []int{2}},
					{a: 0, b: 1, t: 2, neighbors: []int{1}},
				},
			},
			minAnchorings: 1,
		},
		{
			name: "k limit",
			in: input{
				n: 4, p: 1, k: 2, x: 0, m: 5,
				devices: []device{
					{},
					{a: 0, b: 23, t: 1, neighbors: []int{2, 3, 4}},
					{a: 0, b: 23, t: 1, neighbors: []int{1}},
					{a: 0, b: 23, t: 1, neighbors: []int{1}},
					{a: 0, b: 23, t: 1, neighbors: []int{1}},
				},
			},
			minAnchorings: 3,
		},
		{
			name: "ping pong",
			in: input{
				n: 2, p: 1, k: 1, x: 0, m: 6,
				devices: []device{
					{},
					{a: 0, b: 23, t: 1, neighbors: []int{2}},
					{a: 0, b: 23, t: 1, neighbors: []int{1}},
				},
			},
			minAnchorings: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildStrategyV2(tt.in)
			if err := validate(tt.in, got); err != nil {
				t.Fatalf("buildStrategyV2() returned invalid strategy: %v", err)
			}
			if len(got)+1 < tt.minAnchorings {
				t.Fatalf("buildStrategyV2() anchorings = %d, want at least %d", len(got)+1, tt.minAnchorings)
			}
		})
	}
}

func Test_buildStrategyV2_maxSize(t *testing.T) {
	n := 1000
	devices := make([]device, n+1)
	for i := 1; i <= n; i++ {
		devices[i] = device{a: 0, b: 23, t: 1}
		if i > 1 {
			devices[i].neighbors = append(devices[i].neighbors, i-1)
		}
		if i < n {
			devices[i].neighbors = append(devices[i].neighbors, i+1)
		}
	}
	in := input{n: n, p: 1, k: 1, x: 0, m: 1000, devices: devices}
	got := buildStrategyV2(in)
	if err := validate(in, got); err != nil {
		t.Fatalf("buildStrategyV2(maxSize) returned invalid strategy: %v", err)
	}
	if len(got)+1 < 1 {
		t.Fatalf("buildStrategyV2(maxSize) anchorings = %d, want at least 1", len(got)+1)
	}
}
