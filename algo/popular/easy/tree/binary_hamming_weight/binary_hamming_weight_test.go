package main

import (
	"fmt"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				num: uint32(0b00000000000000000000000000001011),
			},
			want: 3,
		},
		{
			name: "CASE-2",
			args: args{
				num: uint32(0b00000000000000000000000010000000),
			},
			want: 1,
		},
		{
			name: "CASE-3",
			args: args{
				num: uint32(0b11111111111111111111111111111101),
			},
			want: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRightShift(t *testing.T) {
	x := uint32(0b00000000000000000000000000001011)

	// Т.е. пока число x ни станет равно 0 (x > uint32(0)),
	// мы будем сдвигать биты этого числа вправо.
	for ; x > uint32(0); x >>= 1 { // блок x>>=1 выполняется после участка внутри цикла
		fmt.Printf("x=%d[%08b] x>>=1 = %d[%08b] \n", x, x, x>>1, x>>1)
	}
}
