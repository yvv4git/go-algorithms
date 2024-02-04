package evaluate_division

import (
	"reflect"
	"testing"
)

func Test_calcEquationV1(t *testing.T) {
	/*
		METHOD: BFS(Breadth-First Search)
		Time complexity: O(n^2)
		Space complexity: O(n)
	*/
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}

	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "CASE-1",
			args: args{
				equations: [][]string{{"a", "b"}, {"b", "c"}},
				values:    []float64{2.0, 3.0},
				queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			},
			want: []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
		},
		{
			name: "CASE-2",
			args: args{
				equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
				values:    []float64{1.5, 2.5, 5.0},
				queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			},
			want: []float64{3.75000, 0.40000, 5.00000, 0.20000},
		},
		{
			name: "CASE-3",
			args: args{
				equations: [][]string{{"a", "b"}},
				values:    []float64{0.5},
				queries:   [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			},
			want: []float64{0.50000, 2.00000, -1.00000, -1.00000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquationV1(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquationV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
