package kth_factor_of_n

import "testing"

func Test_kthFactorV1(t *testing.T) {
	type args struct {
		n int
		k int
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
				n: 12,
				k: 3,
			},
			want: 3,
			desc: "Factors of 12 [1, 2, 3, 4, 6, 12]. k=3 is 3",
		},
		{
			name: "CASE-2",
			args: args{
				n: 7,
				k: 2,
			},
			want: 7,
			desc: "Factors of 7 [1, 7]. k=2 is 7",
		},
		{
			name: "CASE-3",
			args: args{
				n: 4,
				k: 4,
			},
			want: -1,
			desc: "Factors of 4 [1, 2, 4]. k=4. Thera is only 3 factors. We should return -1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthFactorV1(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kthFactorV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
