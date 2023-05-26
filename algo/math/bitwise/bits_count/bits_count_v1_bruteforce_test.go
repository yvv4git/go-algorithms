package bits_count

import "testing"

func Test_bitsCountV1(t *testing.T) {
	type args struct {
		n int8
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
				n: 5,
			},
			want: 2,
			desc: `5(00000101) -> 2`,
		},
		{
			name: "CASE-2",
			args: args{
				n: 10,
			},
			want: 2,
			desc: `10(00001010) -> 2`,
		},
		{
			name: "CASE-3",
			args: args{
				n: 11,
			},
			want: 3,
			desc: `11(00001011) -> 3 ones`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitsCountV1(tt.args.n); got != tt.want {
				t.Errorf("bitsCountV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
