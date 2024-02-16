package main

import "testing"

func TestGraph_DFS(t *testing.T) {
	type fields struct {
		Vertices int
		adjList  map[int][]int
	}

	type args struct {
		start int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Case 1: Linear Graph",
			fields: fields{
				Vertices: 4,
				adjList: map[int][]int{
					0: {1},
					1: {2},
					2: {3},
					3: {},
				},
			},
			args: args{start: 0},
		},
		{
			name: "Test Case 2: Tree Graph",
			fields: fields{
				Vertices: 5,
				adjList: map[int][]int{
					0: {1, 2},
					1: {3, 4},
					2: {},
					3: {},
					4: {},
				},
			},
			args: args{start: 0},
		},
		{
			name: "Test Case 3: Disconnected Graph",
			fields: fields{
				Vertices: 4,
				adjList: map[int][]int{
					0: {1},
					1: {2},
					2: {},
					3: {},
				},
			},
			args: args{start: 0},
		},
		{
			name: "Test Case 4: Single Vertex Graph",
			fields: fields{
				Vertices: 1,
				adjList: map[int][]int{
					0: {},
				},
			},
			args: args{start: 0},
		},
		{
			name: "Test Case 5: Empty Graph",
			fields: fields{
				Vertices: 0,
				adjList:  map[int][]int{},
			},
			args: args{start: 0}, // This test case is expected to fail
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Vertices: tt.fields.Vertices,
				adjList:  tt.fields.adjList,
			}
			g.DFS(tt.args.start)
		})
	}
}
