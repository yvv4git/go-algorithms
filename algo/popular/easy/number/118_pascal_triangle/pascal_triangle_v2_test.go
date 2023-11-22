package _18_pascal_triangle

import (
	"reflect"
	"testing"
)

func Test_generateV2(t *testing.T) {
	type args struct {
		numRows int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "CASE-1",
			args: args{
				numRows: 5,
			},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			name: "CASE-2",
			args: args{
				numRows: 1,
			},
			want: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateV2(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
