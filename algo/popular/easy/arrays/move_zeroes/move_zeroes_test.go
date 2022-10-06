package move_zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}

	testCases := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
