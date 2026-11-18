package main

import "testing"

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

func greedyRouteLen(points []Point) float64 {
	return routeLength(points, orderToRoute(greedyOrder(points)))
}

func Test_annealRoute(t *testing.T) {
	tests := []struct {
		name   string
		points []Point
	}{
		{
			name:   "example 1",
			points: []Point{{0, 0}, {1, 1}, {0, 1}, {1, 0}},
		},
		{
			name: "example 2",
			points: []Point{
				{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10},
				{11, 12}, {13, 14}, {15, 16}, {17, 18}, {19, 20},
			},
		},
		{
			name:   "single village",
			points: []Point{{5, 7}},
		},
		{
			name:   "two villages",
			points: []Point{{0, 0}, {2, 0}},
		},
		{
			name:   "negative coordinates",
			points: []Point{{0, 0}, {-1, 0}, {-2, 0}, {-2, -1}},
		},
		{
			name:   "random-like fixed set",
			points: []Point{{0, 0}, {2, 3}, {5, 1}, {6, 4}, {3, 5}, {1, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			route := annealRoute(tt.points)
			if !isValidRoute(len(tt.points), route) {
				t.Fatalf("invalid route %v for points %v", route, tt.points)
			}
			gotLen := routeLength(tt.points, route)
			greedyLen := greedyRouteLen(tt.points)
			if gotLen > greedyLen+1e-9 {
				t.Fatalf("route length %v worse than greedy %v", gotLen, greedyLen)
			}
		})
	}
}

func Test_annealRoute_maxSize(t *testing.T) {
	const n = 3000
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		points[i] = Point{x: int64(i), y: int64(i % 100)}
	}
	route := annealRoute(points)
	if !isValidRoute(n, route) {
		t.Fatalf("route is not valid for n=%d", n)
	}
}
