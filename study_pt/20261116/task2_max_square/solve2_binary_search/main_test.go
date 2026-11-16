package main

import "testing"

func Test_maxSquareSideV2(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		m    int64
		vert []int64
		hor  []int64
		want int64
	}{
		{name: "example 1", n: 10, m: 10, vert: []int64{3}, hor: []int64{5}, want: 5},
		{name: "example 2", n: 4, m: 7, vert: []int64{}, hor: []int64{}, want: 4},
		{name: "example 3", n: 5, m: 7, vert: []int64{5, 0}, hor: []int64{3, 1, 2}, want: 4},
		{name: "no streets square city", n: 6, m: 6, vert: []int64{}, hor: []int64{}, want: 6},
		{name: "only vertical streets", n: 10, m: 4, vert: []int64{2, 7}, hor: []int64{}, want: 4},
		{name: "only horizontal streets", n: 4, m: 10, vert: []int64{}, hor: []int64{3, 8}, want: 4},
		{name: "streets on borders", n: 10, m: 10, vert: []int64{0, 10}, hor: []int64{0, 10}, want: 10},
		{name: "all streets block everything", n: 5, m: 5, vert: []int64{0, 1, 2, 3, 4, 5}, hor: []int64{0, 1, 2, 3, 4, 5}, want: 1},
		{name: "unsorted street order", n: 10, m: 10, vert: []int64{9, 1, 5}, hor: []int64{8, 2}, want: 4},
		{name: "duplicate coordinates", n: 10, m: 10, vert: []int64{3, 3, 3}, hor: []int64{4, 4}, want: 6},
		{name: "narrow city", n: 1, m: 100, vert: []int64{}, hor: []int64{50}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSquareSideV2(tt.n, tt.m, tt.vert, tt.hor); got != tt.want {
				t.Fatalf("maxSquareSideV2(%d, %d, %v, %v) = %d, want %d",
					tt.n, tt.m, tt.vert, tt.hor, got, tt.want)
			}
		})
	}
}

func Test_maxSquareSideV2_maxSize(t *testing.T) {
	const (
		n     = int64(1000000000)
		m     = int64(1000000000)
		count = 100000
	)
	vert := make([]int64, count)
	hor := make([]int64, count)
	step := n / int64(count+1)
	for i := 0; i < count; i++ {
		vert[i] = step * int64(i+1)
		hor[i] = step * int64(i+1)
	}
	want := n - step*int64(count)
	if got := maxSquareSideV2(n, m, vert, hor); got != want {
		t.Fatalf("maxSquareSideV2(maxSize) = %d, want %d", got, want)
	}
}
