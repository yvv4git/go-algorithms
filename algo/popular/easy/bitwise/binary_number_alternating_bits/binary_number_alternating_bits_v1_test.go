package binary_number_alternating_bits

import "testing"

func Test_hasAlternatingBitsV1(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				n: 5,
			},
			want: true,
			desc: "Explanation: The binary representation of 5 is: 101",
		},
		{
			name: "CASE-2",
			args: args{
				n: 7,
			},
			want: false,
			desc: "Explanation: The binary representation of 7 is: 111.",
		},
		{
			name: "CASE-3",
			args: args{
				n: 11,
			},
			want: false,
			desc: "Explanation: The binary representation of 11 is: 1011.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAlternatingBitsV1(tt.args.n); got != tt.want {
				t.Errorf("hasAlternatingBitsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
