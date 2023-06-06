package binary_gap

import "testing"

func Test_binaryGapV3(t *testing.T) {
	type args struct {
		n int
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
				n: 22,
			},
			want: 2,
			desc: `Explanation: 22 in binary is "10110".
					The first adjacent pair of 1's is "10110" with a distance of 2.
					The second adjacent pair of 1's is "10110" with a distance of 1.
					The answer is the largest of these two distances, which is 2.
					Note that "10110" is not a valid pair since there is a 1 separating the two 1's underlined.`,
		},
		{
			name: "CASE-2",
			args: args{
				n: 8,
			},
			want: 0,
			desc: `Explanation: 8 in binary is "1000".
					There are not any adjacent pairs of 1's in the binary representation of 8, so we return 0.`,
		},
		{
			name: "CASE-3",
			args: args{
				n: 5,
			},
			want: 2,
			desc: `Explanation: 5 in binary is "101".`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGapV3(tt.args.n); got != tt.want {
				t.Errorf("binaryGapV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
