package sort_colors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortColorsV7(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{2, 0, 1},
			},
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColorsV7(tt.args.nums)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
