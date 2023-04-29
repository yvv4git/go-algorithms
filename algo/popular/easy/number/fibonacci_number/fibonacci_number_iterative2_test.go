package fibonacci_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fibonacciNumberIterative2(t *testing.T) {
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
				n: 2,
			},
			want: 1,
		},
		{
			name: "CASE-2",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "CASE-3",
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "CASE-4",
			args: args{
				n: 5,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fibonacciNumberIterative2(tt.args.n)
			//t.Logf("For number=%d result=%d", tt.args.n, result)
			assert.Equal(t, tt.want, result)
		})
	}
}
