package merge_sort_single_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		items []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				items: []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "CASE-2",
			args: args{
				items: []int{10, 6},
			},
			want: []int{6, 10},
		},
		{
			name: "CASE-3",
			args: args{
				items: []int{10},
			},
			want: []int{10},
		},
		{
			name: "CASE-4",
			args: args{
				items: []int{},
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeSort(tt.args.items)
			t.Logf("%#v", result)
			assert.Equal(t, tt.want, result)
		})
	}
}
