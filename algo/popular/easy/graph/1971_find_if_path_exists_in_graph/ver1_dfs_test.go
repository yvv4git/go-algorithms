package _971_find_if_path_exists_in_graph

import "testing"

func Test_validPath(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		start int
		end   int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				n:     3,
				edges: [][]int{{0, 1}, {1, 2}, {2, 0}},
				start: 0,
				end:   2,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
				start: 0,
				end:   5,
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				n:     1,
				edges: [][]int{},
				start: 0,
				end:   0,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPathV1(tt.args.n, tt.args.edges, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("validPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
