package dijkstra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraph_CalcMinPath(t *testing.T) {
	type edge struct {
		origin  string
		destiny string
		weight  float64
	}

	type pair struct {
		from string
		to   string
	}

	type args struct {
		edges  []edge
		target pair
	}

	type want struct {
		pathLen float64
		path    []string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "CASE-1",
			args: args{
				edges: []edge{
					{"BTC", "LTC", 6.0},
					{"USD", "LTC", 6.0},
					{"ETH", "BTC", 3.0},
				},
				target: pair{
					from: "USD",
					to:   "ETH",
				},
			},
			want: want{
				pathLen: 15.0,
				path:    []string{"USD", "LTC", "BTC", "ETH"},
			},
		},
		{
			name: "CASE-2",
			args: args{
				edges: []edge{
					{"BTC", "USD", 5.0},
					{"USD", "LTC", 7.0},
				},
				target: pair{
					from: "BTC",
					to:   "LTC",
				},
			},
			want: want{
				pathLen: 12.0,
				path:    []string{"BTC", "USD", "LTC"},
			},
		},
		{
			name: "CASE-3",
			args: args{
				edges: []edge{
					{"BTC", "USD", 5.3},
					{"ETH", "LTC", 3.5},
				},
				target: pair{
					from: "BTC",
					to:   "LTC",
				},
			},
			want: want{
				pathLen: 0,
				path:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraph()
			for _, e := range tt.args.edges {
				graph.AddEdge(e.origin, e.destiny, e.weight)
			}

			pathLen, path := graph.CalcMinPath(tt.args.target.from, tt.args.target.to)
			assert.Equalf(t, tt.want.pathLen, pathLen, "problems with the length of the path")
			assert.Equalf(t, tt.want.path, path, "problems with path")
		})
	}
}
