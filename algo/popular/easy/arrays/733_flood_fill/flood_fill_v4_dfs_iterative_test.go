package _33_flood_fill

import (
	"reflect"
	"testing"
)

func Test_floodFillV4(t *testing.T) {
	type args struct {
		image [][]int
		sr    int
		sc    int
		color int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "CASE-1",
			args: args{
				image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
				sr:    1,
				sc:    1,
				color: 2,
			},
			want: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			name: "CASE-2",
			args: args{
				image: [][]int{{0, 0, 0}, {0, 0, 0}},
				sr:    0,
				sc:    0,
				color: 0,
			},
			want: [][]int{{0, 0, 0}, {0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFillV4(tt.args.image, tt.args.sr, tt.args.sc, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFillV4() = %v, want %v", got, tt.want)
			}
		})
	}
}
