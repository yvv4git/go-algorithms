package main

import "testing"

func completeEdgesV3(n int) [][2]int {
	edges := make([][2]int, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, [2]int{i, j})
		}
	}
	return edges
}

func Test_hasIndependentSetV3(t *testing.T) {
	twoCliques := [][2]int{
		{0, 1}, {1, 2}, {0, 2},
		{3, 4}, {4, 5}, {3, 5},
	}
	tests := []struct {
		name     string
		vertices int
		target   int
		edges    [][2]int
		want     bool
	}{
		{
			name:     "example 1",
			vertices: 5,
			target:   3,
			edges:    [][2]int{{0, 1}, {2, 3}, {3, 4}},
			want:     true,
		},
		{
			name:     "example 2",
			vertices: 5,
			target:   4,
			edges:    [][2]int{{0, 1}, {2, 3}, {3, 4}},
			want:     false,
		},
		{
			name:     "single vertex",
			vertices: 1,
			target:   1,
			edges:    nil,
			want:     true,
		},
		{
			name:     "target greater than vertices",
			vertices: 1,
			target:   2,
			edges:    nil,
			want:     false,
		},
		{
			name:     "target zero",
			vertices: 3,
			target:   0,
			edges:    [][2]int{{0, 1}, {1, 2}},
			want:     true,
		},
		{
			name:     "all vertices without edges",
			vertices: 5,
			target:   5,
			edges:    nil,
			want:     true,
		},
		{
			name:     "complete triangle size two",
			vertices: 3,
			target:   2,
			edges:    completeEdgesV3(3),
			want:     false,
		},
		{
			name:     "duplicate edges",
			vertices: 3,
			target:   2,
			edges:    [][2]int{{0, 1}, {0, 1}},
			want:     true,
		},
		{
			name:     "star leaves",
			vertices: 5,
			target:   4,
			edges:    [][2]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			want:     true,
		},
		{
			name:     "star all vertices",
			vertices: 5,
			target:   5,
			edges:    [][2]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
			want:     false,
		},
		{
			name:     "path distant vertices",
			vertices: 4,
			target:   2,
			edges:    [][2]int{{0, 1}, {1, 2}, {2, 3}},
			want:     true,
		},
		{
			name:     "two triangles one from each",
			vertices: 6,
			target:   2,
			edges:    twoCliques,
			want:     true,
		},
		{
			name:     "two triangles three vertices",
			vertices: 6,
			target:   3,
			edges:    twoCliques,
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasIndependentSetV3(tt.vertices, tt.target, tt.edges); got != tt.want {
				t.Fatalf("hasIndependentSetV3(%d, %d, %v) = %v, want %v", tt.vertices, tt.target, tt.edges, got, tt.want)
			}
		})
	}
}

func Test_hasIndependentSetV3_maxSize(t *testing.T) {
	const vertices = 50
	if !hasIndependentSetV3(vertices, vertices, nil) {
		t.Fatalf("hasIndependentSetV3(%d, %d, nil) = false, want true", vertices, vertices)
	}
}
