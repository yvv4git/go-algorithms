package tree

import "testing"

func TestTree_CalculateHeight(t1 *testing.T) {
	type fields struct {
		nodes map[int64]*Node
		root  *Node
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test case with a single node",
			fields: fields{
				nodes: map[int64]*Node{
					1: {ID: 1, Name: "Root", ParentID: 0},
				},
				root: &Node{ID: 1, Name: "Root", ParentID: 0},
			},
			want: 1,
		},
		{
			name: "Test case with a balanced tree",
			fields: fields{
				nodes: map[int64]*Node{
					1: {ID: 1, Name: "Root", ParentID: 0},
					2: {ID: 2, Name: "Child 1", ParentID: 1},
					3: {ID: 3, Name: "Child 2", ParentID: 1},
					4: {ID: 4, Name: "Grandchild 1", ParentID: 2},
					5: {ID: 5, Name: "Grandchild 2", ParentID: 3},
				},
				root: &Node{ID: 1, Name: "Root", ParentID: 0},
			},
			want: 3,
		},
		{
			name: "Test case with an unbalanced tree",
			fields: fields{
				nodes: map[int64]*Node{
					1: {ID: 1, Name: "Root", ParentID: 0},
					2: {ID: 2, Name: "Child 1", ParentID: 1},
					3: {ID: 3, Name: "Child 2", ParentID: 2},
					4: {ID: 4, Name: "Child 3", ParentID: 3},
				},
				root: &Node{ID: 1, Name: "Root", ParentID: 0},
			},
			want: 4,
		},
		{
			name: "Test case with an empty tree",
			fields: fields{
				nodes: map[int64]*Node{},
				root:  nil,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				nodes: tt.fields.nodes,
				root:  tt.fields.root,
			}
			if got := t.CalculateHeight(); got != tt.want {
				t1.Errorf("CalculateHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Print(t1 *testing.T) {
	type fields struct {
		nodes map[int64]*Node
		root  *Node
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test case with a single node",
			fields: fields{
				nodes: map[int64]*Node{
					1: {ID: 1, Name: "Root", ParentID: 0},
				},
				root: &Node{ID: 1, Name: "Root", ParentID: 0},
			},
		},
		{
			name: "Test case with a balanced tree",
			fields: fields{
				nodes: map[int64]*Node{
					1: {ID: 1, Name: "Root", ParentID: 0},
					2: {ID: 2, Name: "Child 1", ParentID: 1},
					3: {ID: 3, Name: "Child 2", ParentID: 1},
					4: {ID: 4, Name: "Grandchild 1", ParentID: 2},
					5: {ID: 5, Name: "Grandchild 2", ParentID: 3},
				},
				root: &Node{ID: 1, Name: "Root", ParentID: 0},
			},
		},
		{
			name: "Test case with an unbalanced tree",
			fields: fields{
				nodes: map[int64]*Node{
					1: {ID: 1, Name: "Root", ParentID: 0},
					2: {ID: 2, Name: "Child 1", ParentID: 1},
					3: {ID: 3, Name: "Child 2", ParentID: 2},
					4: {ID: 4, Name: "Child 3", ParentID: 3},
				},
				root: &Node{ID: 1, Name: "Root", ParentID: 0},
			},
		},
		{
			name: "Test case with a deeply branched tree",
			fields: fields{
				nodes: map[int64]*Node{
					1:  {ID: 1, Name: "Root", ParentID: 0},
					2:  {ID: 2, Name: "Child 1", ParentID: 1},
					3:  {ID: 3, Name: "Child 2", ParentID: 1},
					4:  {ID: 4, Name: "Child 3", ParentID: 1},
					5:  {ID: 5, Name: "Grandchild 1", ParentID: 2},
					6:  {ID: 6, Name: "Grandchild 2", ParentID: 2},
					7:  {ID: 7, Name: "Grandchild 3", ParentID: 3},
					8:  {ID: 8, Name: "Grandchild 4", ParentID: 3},
					9:  {ID: 9, Name: "Grandchild 5", ParentID: 4},
					10: {ID: 10, Name: "Grandchild 6", ParentID: 4},
				},
				root: &Node{ID: 1, Name: "Root", ParentID: 0},
			},
		},
		{
			name: "Test case with an empty tree",
			fields: fields{
				nodes: map[int64]*Node{},
				root:  nil,
			},
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tree{
				nodes: tt.fields.nodes,
				root:  tt.fields.root,
			}
			t.Print()
		})
	}
}
