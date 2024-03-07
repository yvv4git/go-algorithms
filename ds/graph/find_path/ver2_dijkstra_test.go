package main

import "testing"

func Test_validPathV2(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		//{
		//	name: "Test Case 1: Valid path exists",
		//	args: args{
		//		n:           3,
		//		edges:       [][]int{{0, 1}, {1, 2}, {2, 0}},
		//		source:      0,
		//		destination: 2,
		//	},
		//	want: true,
		//},
		{
			name: "Test Case 2: Valid path does not exist",
			args: args{
				n:           6,
				edges:       [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
				source:      0,
				destination: 5,
			},
			want: false,
		},
		{
			name: "Test Case 3: Single node graph",
			args: args{
				n:           1,
				edges:       [][]int{},
				source:      0,
				destination: 0,
			},
			want: true,
		},
		{
			name: "Test Case 4: Disconnected graph",
			args: args{
				n:           4,
				edges:       [][]int{{0, 1}, {2, 3}},
				source:      0,
				destination: 3,
			},
			want: false,
		},
		//{
		//	name: "Test Case 5: Large graph with valid path",
		//	args: args{
		//		n:           1000,
		//		edges:       generateLargeGraph(1000, 999),
		//		source:      0,
		//		destination: 999,
		//	},
		//	want: true,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPathV2(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination); got != tt.want {
				t.Errorf("validPathV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
