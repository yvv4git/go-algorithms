package main

import "testing"

func Test_greedyRoute(t *testing.T) {
	tests := []struct {
		name   string
		points []Point
		want   []int
	}{
		{
			name:   "example 1",
			points: []Point{{0, 0}, {1, 1}, {0, 1}, {1, 0}},
			want:   []int{1, 3, 2, 4, 1},
		},
		{
			name: "example 2",
			points: []Point{
				{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10},
				{11, 12}, {13, 14}, {15, 16}, {17, 18}, {19, 20},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1},
		},
		{
			name:   "single village",
			points: []Point{{5, 7}},
			want:   []int{1, 1},
		},
		{
			name:   "two villages",
			points: []Point{{0, 0}, {2, 0}},
			want:   []int{1, 2, 1},
		},
		{
			name:   "tie chooses smaller number",
			points: []Point{{0, 0}, {1, 0}, {0, 1}},
			want:   []int{1, 2, 3, 1},
		},
		{
			name:   "negative coordinates",
			points: []Point{{0, 0}, {-1, 0}, {-2, 0}},
			want:   []int{1, 2, 3, 1},
		},
		{
			name:   "collinear points",
			points: []Point{{0, 0}, {3, 0}, {1, 0}, {2, 0}},
			want:   []int{1, 3, 4, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := greedyRoute(tt.points)
			if len(got) != len(tt.want) {
				t.Fatalf("greedyRoute(%v) = %v, want %v", tt.points, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("greedyRoute(%v) = %v, want %v", tt.points, got, tt.want)
				}
			}
		})
	}
}

func isValidRoute(n int, route []int) bool {
	if len(route) != n+1 {
		return false
	}
	if route[0] != 1 || route[n] != 1 {
		return false
	}
	seen := make([]bool, n+1)
	for i := 0; i < n; i++ {
		v := route[i]
		if v < 1 || v > n || seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}

func Test_greedyRoute_maxSize(t *testing.T) {
	const n = 3000
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		points[i] = Point{x: int64(i), y: int64(i % 100)}
	}
	route := greedyRoute(points)
	if !isValidRoute(n, route) {
		t.Fatalf("route is not valid for n=%d", n)
	}
}
