package main

import (
	"reflect"
	"testing"
)

func Test_allCellsDistOrder(t *testing.T) {
	type args struct {
		rows    int
		cols    int
		rCenter int
		cCenter int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1x1 matrix with center at (0,0)",
			args: args{rows: 1, cols: 1, rCenter: 0, cCenter: 0},
			want: [][]int{{0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCellsDistOrder(tt.args.rows, tt.args.cols, tt.args.rCenter, tt.args.cCenter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allCellsDistOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
