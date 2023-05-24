package number_complement

import (
	"testing"
)

func Test_findComplementV1(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				num: 5,
			},
			want: 2,
			desc: "The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.",
		},
		{
			name: "CASE-2",
			args: args{
				num: 1,
			},
			want: 0,
			desc: "The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findComplementV1(tt.args.num); got != tt.want {
				t.Errorf("findComplementV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
