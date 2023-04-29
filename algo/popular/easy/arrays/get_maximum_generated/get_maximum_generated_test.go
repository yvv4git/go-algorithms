package get_maximum_generated

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getMaximumGenerated(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				n: 7,
			},
			want: 3,
		},
		{
			name: "CASE-2",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "CASE-3",
			args: args{
				n: 3,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getMaximumGenerated(tt.args.n))
		})
	}
}
