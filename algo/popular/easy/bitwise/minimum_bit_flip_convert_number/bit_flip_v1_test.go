package minimum_bit_flip_convert_number

import "testing"

func Test_minBitFlipsV1(t *testing.T) {
	type args struct {
		start int
		goal  int
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
				start: 10,
				goal:  7,
			},
			want: 3,
			desc: `The binary representation of 10 and 7 are 1010 and 0111 respectively. We can convert 10 to 7 in 3 steps:
					- Flip the first bit from the right: 1010 -> 1011.
					- Flip the third bit from the right: 1011 -> 1111.
					- Flip the fourth bit from the right: 1111 -> 0111.
					It can be shown we cannot convert 10 to 7 in less than 3 steps. Hence, we return 3.`,
		},
		{
			name: "CASE-2",
			args: args{
				start: 3,
				goal:  4,
			},
			want: 3,
			desc: `The binary representation of 3 and 4 are 011 and 100 respectively. We can convert 3 to 4 in 3 steps:
					- Flip the first bit from the right: 011 -> 010.
					- Flip the second bit from the right: 010 -> 000.
					- Flip the third bit from the right: 000 -> 100.
					It can be shown we cannot convert 3 to 4 in less than 3 steps. Hence, we return 3.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minBitFlipsV1(tt.args.start, tt.args.goal); got != tt.want {
				t.Errorf("minBitFlipsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
