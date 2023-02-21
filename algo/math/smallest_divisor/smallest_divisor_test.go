package smallest_divisor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallestDivisor(t *testing.T) {
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
				n: 5,
			},
			want: 5,
		},
		{
			name: "CASE-2",
			args: args{
				n: 25,
			},
			want: 5,
		},
		{
			name: "CASE-3",
			args: args{
				n: 64,
			},
			want: 2,
		},
		{
			name: "CASE-4",
			args: args{
				n: 30,
			},
			want: 2,
		},
		{
			name: "CASE-5",
			args: args{
				n: 31,
			},
			want: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SmallestDivisor(tt.args.n)
			assert.Equal(t, tt.want, result)

		})
	}
}
