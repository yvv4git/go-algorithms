package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		a []int
		b []int
	}

	tests := []struct {
		name string
		args args
		want []int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				a: []int{},
				b: []int{},
			},
			want: nil,
			desc: "Когда в исходных массивах пусто",
		},
		{
			name: "CASE-2",
			args: args{
				a: []int{1, 2, 3},
				b: []int{3, 4, 5},
			},
			want: []int{3},
			desc: "Когда есть пересекающиеся элементы",
		},
		{
			name: "CASE-3",
			args: args{
				a: []int{1, 2},
				b: []int{3, 4, 5},
			},
			want: nil,
			desc: "Когда нет пересекающихся элементов",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := intersection(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, result)
		})
	}
}
