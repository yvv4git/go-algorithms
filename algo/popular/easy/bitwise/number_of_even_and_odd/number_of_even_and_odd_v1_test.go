package number_of_even_and_odd

import (
	"reflect"
	"testing"
)

func Test_evenOddBitV1(t *testing.T) {
	type args struct {
		n int
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
				n: 17,
			},
			want: []int{2, 0},
			desc: `Explanation: The binary representation of 17 is 10001. 
					It contains 1 on the 0th and 4th indices. 
					There are 2 even and 0 odd indices.`,
		},
		{
			name: "CASE-2",
			args: args{
				n: 2,
			},
			want: []int{0, 1},
			desc: `Explanation: The binary representation of 2 is 10.
					It contains 1 on the 1st index. 
					There are 0 even and 1 odd indices.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenOddBitV1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evenOddBitV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
