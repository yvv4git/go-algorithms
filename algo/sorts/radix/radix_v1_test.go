package radix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_radixSortV1(t *testing.T) {
	type args struct {
		list []int32
	}

	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "CASE-1",
			args: args{
				list: []int32{91, 28, 73, 46, 55, 64, 37, 82, 19},
			},
			want: []int32{19, 28, 37, 46, 55, 64, 73, 82, 91},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			radixSortV1(tt.args.list)
			assert.Equal(t, tt.want, tt.args.list)
		})
	}
}
